package main

import (
	"sync"
	"fmt"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	
	"local.com/13sai/game/config"
	"local.com/13sai/game/web"
	"local.com/13sai/game/services"
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

	// pingServer()
	// return

	num := viper.GetInt("addr_num")
	port := viper.GetInt("addr")
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		web := web.StartHttp()
		go func (i int) {
			addr := fmt.Sprintf(":%d", i+port)
			services.Register(web, addr)
			// s.Run()
		}(i)
		// time.Sleep(time.Second*10);
	}

	wg.Wait()
}

func pingServer() {
	time.NewTicker(time.Second)
    for i := 0; i < 10; i++ {
		services.Discover()
    }
}