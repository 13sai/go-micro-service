package services

import (
	"fmt"
	"math/rand"
	"context"

	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/service"
	gserver "github.com/micro/go-micro/v2/service/grpc"

    "github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/consul/v2"
	// "github.com/micro/go-micro/v2/server"


	// "github.com/micro/go-micro/server/grpc"
    // "google.golang.org/grpc"
	"local.com/13sai/microService/hello"
	// gDemo "local.com/13sai/microService/grpc"
)

// 注册服务
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

type Server struct {
    // hello.DemoService
}

func (t *Server) SayHello(ctx context.Context, req *hello.HelloRequest, rsp *hello.HelloReply) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func RegisterGRpc(addr string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := consul.NewRegistry(
        registry.Addrs(viper.GetString("service_addr")),
	)

	// create GRPC service
	s := gserver.NewService(
		service.Name(viper.GetString("service_name")),
		service.Registry(r),
		service.AfterStart(func() error {
			return nil
		}),
		service.Context(ctx),
	)
	fmt.Println(addr)

	// register test handler
	hello.RegisterDemoHandler(s.Server(), &Server{})
	s.Run()
}