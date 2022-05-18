package main

import (
	"gateway/reverse_proxy"
	"gateway/reverse_proxy/load_balance"
	"log"
	"net/http"
)

func main() {
	//// 从配置文件读取配置
	//conf.Init()
	//
	//// 装载路由
	//r := server.NewRouter()
	//r.Run(":8880")

	mConf, err := load_balance.NewLoadBalanceCheckConf(
		"http://%s/base",
		map[string]string{"127.0.0.1:2003": "20",
			"127.0.0.1:2004": "20",
			"127.0.0.1:2001": "40"})
	if err != nil {
		panic(err)
	}
	rb := load_balance.LoadBalanceFactoryWithConf(load_balance.LbRoundRobin, mConf)
	proxy := reverse_proxy.NewLoadBalanceReverseProxy(rb)
	log.Fatal(http.ListenAndServe("127.0.0.1:2010", proxy))
}
