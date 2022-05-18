package load_balance

import (
	"fmt"
	"net"
	"reflect"
	"sort"
	"time"
)

const (
	DefaultCheckMethod    = 0
	DefaultCheckTimeout   = 3  // 心跳检测最大超时时间
	DefaultCheckMaxErrNum = 2  // 心跳检测最大重试次数
	DefaultCheckInterval  = 10 // 心跳检测的间隔时间
)

type LoadBalanceCheckConf struct {
	observers    []Observer        // 观察者集合
	confIpWeight map[string]string // 自定义的负载均衡结点以及其权值
	activeList   []string          // 存活的结点
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

// GetConf 从观察者模式中获取服务器配置的列表
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

// WatchConf 开启一个携程用top连接去进行心跳检测，最大重连次数是2次，每5秒进行一次心跳检测，tcp的超时时间是5秒
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
				} else {
					if _, ok := confIpErrNum[item]; ok {
						confIpErrNum[item] += 1
					} else {
						confIpErrNum[item] = 1
					}
				}
				// 最大重连次数
				if confIpErrNum[item] < DefaultCheckMaxErrNum {
					changedList = append(changedList, item)
				}
			}
			sort.Strings(changedList)
			sort.Strings(s.activeList)
			if !reflect.DeepEqual(changedList, s.activeList) {
				s.UpdateConf(changedList)
			}
			// 每5秒进行心跳检测
			time.Sleep(time.Duration(DefaultCheckInterval) * time.Second)
		}
	}()
}

// UpdateConf 负载均衡结点发生变化时，通知load_balance 的结点进行更新
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
