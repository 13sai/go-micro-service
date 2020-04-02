package services

import (
    "fmt"

	"github.com/micro/go-micro/client/selector"
    "github.com/micro/go-micro/registry"
    "github.com/micro/go-plugins/registry/consul"
	"github.com/spf13/viper"
)

func Discover() {
    consulReg := consul.NewRegistry(
        registry.Addrs(viper.GetString("service_addr")),
    )
    nodes, err := consulReg.GetService(viper.GetString("service_name")) //使用服务名获取服务
    if err != nil {
        fmt.Println(err)
    }
	next := selector.Random(nodes)
    node, err := next()          
    if err != nil {
		fmt.Println(err)
	}
	//可以看到我们的id address还有metadata
    fmt.Println(fmt.Sprintf("id:%s,address:%s", node.Id, node.Address)) 
}