package http_proxy_middleware

import (
	"encoding/json"
	"fmt"
	"gateway/dao"
	"gateway/reverse_proxy/load_balance"
	"gateway/service/model"
	"gateway/util"
	"net"
	"net/http"
	"sync"
	"time"
)

var LoadBalancerHandler *LoadBalancer

type LoadBalancer struct {
	LoadBalanceMap   map[string]load_balance.LoadBalance
	LoadBalanceSlice []load_balance.LoadBalance
	Locker           sync.RWMutex
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		LoadBalanceMap:   map[string]load_balance.LoadBalance{},
		LoadBalanceSlice: []load_balance.LoadBalance{},
		Locker:           sync.RWMutex{},
	}
}

var TransportHandler *Transportor

type Transportor struct {
	TransportMap   map[string]TransportItem
	TransportSlice []*TransportItem
	Locker         sync.RWMutex
}

type TransportItem struct {
	Trans       *http.Transport
	ServiceName string
}

func NewTransportor() *Transportor {
	return &Transportor{
		TransportMap:   map[string]TransportItem{},
		TransportSlice: []*TransportItem{},
		Locker:         sync.RWMutex{},
	}
}

func init() {
	LoadBalancerHandler = NewLoadBalancer()
}

func (l *LoadBalancer) GetLoadBalance(gatewayService *dao.GatewayService) (load_balance.LoadBalance, error) {
	httpRule := model.GatewayServiceHttpRuleVO{}
	err := json.Unmarshal([]byte(gatewayService.RemoteParams), &httpRule)
	if err != nil {
		return nil, err
	}
	schema := "http"
	if httpRule.NeedHttps == 1 {
		schema = "https"
	}
	prefix := ""
	if httpRule.RuleType == util.HTTPRuleTypePrefixURL {
		prefix = httpRule.Rule
	}
	ipList := gatewayService.GetIpListByModel()
	weightList := gatewayService.GetWeightListByModel()

	ipConf := map[string]string{}
	for ipIndex, ipItem := range ipList {
		ipConf[ipItem] = weightList[ipIndex]
	}

	mConf, err := load_balance.NewLoadBalanceCheckConf(
		fmt.Sprintf("%s://%s%s", schema, prefix),
		ipConf,
	)
	if err != nil {
		return nil, err
	}
	return load_balance.LoadBalanceFactoryWithConf(
		load_balance.LbRandom, mConf), nil
}

func (t *Transportor) GetTrans(gatewayService *dao.GatewayService) (*http.Transport, error) {
	httpRule := model.GatewayServiceHttpRuleVO{}
	err := json.Unmarshal([]byte(gatewayService.RemoteParams), &httpRule)
	if err != nil {
		return nil, err
	}

	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: time.Duration(httpRule.UpstreamConnectTimeout),
		}).DialContext,
		MaxIdleConns:          httpRule.UpstreamMaxIDle,
		IdleConnTimeout:       time.Duration(httpRule.UpstreamIDleTimeout),
		ResponseHeaderTimeout: time.Duration(httpRule.UpstreamHeaderTimeout),
	}
	t.TransportSlice = append(t.TransportSlice, &TransportItem{
		Trans:       trans,
		ServiceName: gatewayService.ServiceName,
	})
	return trans, nil
}
