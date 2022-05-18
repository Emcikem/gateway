package load_balance

import (
	"fmt"
	"net"
	"reflect"
	"sort"
	"time"
)

const (
	// default check setting
	DefaultCheckMethod    = 0
	DefaultCheckTimeout   = 5
	DefaultCheckMaxErrNum = 2
	DefaultCheckInterval  = 5
)

type LoadBalanceCheckConf struct {
	observers    []Observer // 观察者集合
	confIpWeight map[string]string
	activeList   []string
	format       string
}

func (s *LoadBalanceCheckConf) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *LoadBalanceCheckConf) NotifyAllObserver() {
	for _, obs := range s.observers {
		obs.Update()
	}
}

// GetConf 获取服务器配置的列表
func (s *LoadBalanceCheckConf) GetConf() []string {
	var confList []string
	for _, ip := range s.activeList {
		weight, ok := s.confIpWeight[ip]
		if !ok {
			weight = "50" // 默认weight
		}
		confList = append(confList, fmt.Sprintf(s.format, ip)+weight)
	}
	return confList
}

func (s *LoadBalanceCheckConf) WatchConf() {
	go func() {
		confIpErrNum := map[string]int{}
		for {
			var changedList []string
			for item, _ := range s.confIpWeight {
				conn, err := net.DialTimeout("tcp", item, time.Duration(DefaultCheckTimeout)*time.Second)
				if err == nil {
					conn.Close()
					if _, ok := confIpErrNum[item]; ok {
						confIpErrNum[item] = 0
					}
				}
				if err != nil {
					if _, ok := confIpErrNum[item]; ok {
						confIpErrNum[item] += 1
					} else {
						confIpErrNum[item] = 1
					}
				}
				if confIpErrNum[item] < DefaultCheckMaxErrNum {
					changedList = append(changedList, item)
				}
			}
			sort.Strings(changedList)
			sort.Strings(s.activeList)
			if !reflect.DeepEqual(changedList, s.activeList) {
				s.UpdateConf(changedList)
			}
			time.Sleep(time.Duration(DefaultCheckInterval) * time.Second)
		}
	}()
}

// UpdateConf 更新配置时，通知监听者也更新
func (s *LoadBalanceCheckConf) UpdateConf(conf []string) {
	s.activeList = conf
	for _, obs := range s.observers {
		obs.Update()
	}
}

func NewLoadBalanceCheckConf(format string, conf map[string]string) (*LoadBalanceCheckConf, error) {
	var aList []string
	// 默认初始化
	for item, _ := range conf {
		aList = append(aList, item)
	}
	mConf := &LoadBalanceCheckConf{format: format, activeList: aList, confIpWeight: conf}
	mConf.WatchConf()
	return mConf, nil
}
