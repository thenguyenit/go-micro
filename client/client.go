package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	proto "github.com/thenguyenit/go-micro/srv/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := grpc.NewService(micro.Name("greeter.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}
