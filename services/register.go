package services

import (
	"fmt"
	"math/rand"

	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
)

func Register(ginRouter *gin.Engine, addr string) {
	//新建一个consul注册的地址，也就是我们consul服务启动的机器ip+端口
    consulReg := consul.NewRegistry(
        registry.Addrs(viper.GetString("service_addr")),
    )

	a := rand.Intn(99999)

	server := web.NewService( //go-micro很灵性的实现了注册和反注册，我们启动后直接ctrl+c退出这个server，它会自动帮我们实现反注册
        web.Name(viper.GetString("service_name")), //注册进consul服务中的service名字
        web.Address(addr), //注册进consul服务中的端口
        web.Handler(ginRouter), //web.Handler()返回一个Option，我们直接把ginRouter穿进去，就可以和gin完美的结合
		web.Registry(consulReg),//注册到哪个服务器伤的consul中
		web.Id(fmt.Sprintf("%d-%s", a, addr)),
	)
	fmt.Println(a)
	// return server
    server.Run()
	// fmt.Println(addr)
}