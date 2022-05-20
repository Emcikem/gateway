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
	LoadBalanceMap   map[string]*LoadBalancerItem
	LoadBalanceSlice []*LoadBalancerItem
	Locker           sync.RWMutex
}

type LoadBalancerItem struct {
	LoadBalance load_balance.LoadBalance
	ServiceName string
}

func NewLoadBalancer() *LoadBalancer {
	return &LoadBalancer{
		LoadBalanceMap:   map[string]*LoadBalancerItem{},
		LoadBalanceSlice: []*LoadBalancerItem{},
		Locker:           sync.RWMutex{},
	}
}

var TransportHandler *Transportor

type Transportor struct {
	TransportMap   map[string]*TransportItem
	TransportSlice []*TransportItem
	Locker         sync.RWMutex
}

type TransportItem struct {
	Trans       *http.Transport
	ServiceName string
}

func NewTransportor() *Transportor {
	return &Transportor{
		TransportMap:   map[string]*TransportItem{},
		TransportSlice: []*TransportItem{},
		Locker:         sync.RWMutex{},
	}
}

func init() {
	LoadBalancerHandler = NewLoadBalancer() // 单例的负载均衡池
	TransportHandler = NewTransportor()     // 单例的transport池
}

func (lbr *LoadBalancer) GetLoadBalance(gatewayService *dao.GatewayService) (load_balance.LoadBalance, error) {
	for _, lbrItem := range lbr.LoadBalanceSlice {
		if lbrItem.ServiceName == gatewayService.ServiceName {
			return lbrItem.LoadBalance, nil
		}
	}

	httpRule := model.GatewayServiceHttpRuleVO{}
	err := json.Unmarshal([]byte(gatewayService.RemoteParams), &httpRule)
	if err != nil {
		return nil, err
	}
	schema := "http"
	if httpRule.NeedHttps == 1 { // 支持https
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

	// 构建服务发现版负载均衡
	mConf, err := load_balance.NewLoadBalanceCheckConf(
		fmt.Sprintf("%s://%s%s", schema, prefix),
		ipConf,
	)
	if err != nil {
		return nil, err
	}
	lb := load_balance.LoadBalanceFactoryWithConf(
		load_balance.LbType(gatewayService.RoundType),
		mConf)

	// 维护单例的负载均衡池
	lbItem := &LoadBalancerItem{
		LoadBalance: lb,
		ServiceName: gatewayService.ServiceName,
	}
	lbr.LoadBalanceSlice = append(lbr.LoadBalanceSlice, lbItem)
	lbr.Locker.Lock()
	defer lbr.Locker.Unlock()
	lbr.LoadBalanceMap[gatewayService.ServiceName] = lbItem
	return lb, nil
}

func (t *Transportor) GetTrans(gatewayService *dao.GatewayService) (*http.Transport, error) {
	httpRule := model.GatewayServiceHttpRuleVO{}
	err := json.Unmarshal([]byte(gatewayService.RemoteParams), &httpRule)
	if err != nil {
		return nil, err
	}
	for _, transItem := range t.TransportSlice {
		if transItem.ServiceName == gatewayService.ServiceName {
			return transItem.Trans, nil
		}
	}
	trans := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: time.Duration(httpRule.UpstreamConnectTimeout) * time.Second,
		}).DialContext,
		MaxIdleConns:          httpRule.UpstreamMaxIDle,
		IdleConnTimeout:       time.Duration(httpRule.UpstreamIDleTimeout) * time.Second,
		ResponseHeaderTimeout: time.Duration(httpRule.UpstreamHeaderTimeout) * time.Second,
	}
	transItem := &TransportItem{
		Trans:       trans,
		ServiceName: gatewayService.ServiceName,
	}
	t.TransportSlice = append(t.TransportSlice, transItem)
	t.Locker.Lock()
	defer t.Locker.Unlock()
	t.TransportMap[gatewayService.ServiceName] = transItem
	return trans, nil
}
