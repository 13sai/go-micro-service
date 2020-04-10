package main

import (
    "context"
    "log"
    "os"
    "time"

    // "google.golang.org/grpc"
	// client "github.com/micro/go-micro/v2/client"
	"github.com/spf13/viper"
	"github.com/spf13/pflag"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-micro/v2/registry"

	"local.com/13sai/microService/hello"

	"local.com/13sai/microService/config"
	// "github.com/micro/go-micro/v2/service"
)

const (
    address     = "127.0.0.1:9140"
    defaultName = "13sai"
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
	reg := consul.NewRegistry(
        registry.Addrs(viper.GetString("service_addr")),
    )

	service := micro.NewService(micro.Registry(reg), micro.Name("greeter.client"))


    // Set up a connection to the server.
	test := hello.NewDemoService(viper.GetString("service_name"), service.Client())

	name := defaultName
    if len(os.Args) > 1 {
        name = os.Args[1]
    }

	// call service
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := test.SayHello(ctx, &hello.HelloRequest{Name: name})
    if err != nil {
        log.Fatalf("could not say hello: %v", err)
    }
    log.Printf("Greeting: %s", r.GetMessage())
}