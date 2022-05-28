package main

import (
	"gateway/conf"
	"gateway/real_service"
	"gateway/router"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	go func() {
		real_service.RunRealService()
	}()

	// 装载路由
	r := router.NewRouter()
	r.Run(":8880")

	//mConf, err := load_balance.NewLoadBalanceCheckConf(
	//	"http://%s/base",
	//	map[string]string{
	//		"127.0.0.1:2000": "10",
	//		"127.0.0.1:2001": "10",
	//		"127.0.0.1:2002": "10",
	//		"127.0.0.1:2003": "10",
	//		"127.0.0.1:2004": "10",
	//		"127.0.0.1:2005": "10",
	//		"127.0.0.1:2006": "10",
	//		"127.0.0.1:2007": "10",
	//		"127.0.0.1:2008": "10",
	//		"127.0.0.1:2009": "10",
	//	})
	//if err != nil {
	//	panic(err)
	//}
	//rb := load_balance.LoadBalanceFactoryWithConf(load_balance.LbRoundRobin, mConf)
	//proxy := reverse_proxy.NewLoadBalanceReverseProxy(nil, rb, nil)
	//log.Fatal(http.ListenAndServe("127.0.0.1:2010", proxy))
}
