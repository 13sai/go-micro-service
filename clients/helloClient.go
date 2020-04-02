package main

import (
    "context"
    "fmt"

    proto "winmicro/proto"

    micro "github.com/micro/go-micro"
)

func main() {
    service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
    service.Init()
    helloservice := proto.NewHelloService("hellooo", service.Client())
    res, err := helloservice.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(res.Msg)
}
