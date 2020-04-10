package main

import (
	"sync"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	
	"local.com/13sai/microService/config"
	// "local.com/13sai/microService/web"
	"local.com/13sai/microService/services"
    // "local.com/13sai/microService/grpc"
)

var (
	cfg = pflag.StringP("config", "c", "", "config file path.")
)

func main() {
	// 初始化配置
	pflag.Parse()
	if err := config.Init(*cfg); err != nil {
	
		panic(err)
	}

	grpcServer()

	httpServer()
}

func grpcServer() {
	num := viper.GetInt("addr_num")
	port := viper.GetInt("rpc_addr")

	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func (i int) {
			services.RegisterGRpc(fmt.Sprintf(":%d",i+port))
		}(i)
	}


	wg.Wait()

	// pingServer()

	
}

func httpServer() {
	num := viper.GetInt("addr_num")
	// port := viper.GetInt("addr")
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		// web := web.StartHttp()
		go func (i int) {
			// addr := fmt.Sprintf(":%d", i+port)
			// services.Register(web, addr)
		}(i)
	}

	pingServer()

	wg.Wait()
}

func pingServer() {
	time.NewTicker(time.Second)
    for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		services.Discover()
    }
}