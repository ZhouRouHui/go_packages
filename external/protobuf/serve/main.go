package main

import (
	"context"
	"fmt"
	"go_packages/external/protobuf/pb"
	"google.golang.org/grpc"
	"net"
)

type HelloAchieve struct {
}

func (h *HelloAchieve) HelloWorld(ctx context.Context, params *pb.Params) (*pb.Resp, error) {
	return &pb.Resp{
		Age: 18,
	}, nil
}

func main() {
	/**
	proto 文件定义好 service，里面有 rpc 服务的代码，编译后 go 格式的代码在 pb.go 中
	这里需要注册 rpc 服务才能使用
	*/

	// 初始一个 grpc 对象
	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterHelloServer(grpcServer, new(HelloAchieve))

	// 设置监听， 指定 IP、port
	listener, err := net.Listen("tcp", ":9002")
	if err != nil {
		fmt.Println("get listener failed,", err)
		return
	}
	defer listener.Close()

	// 启动服务。---- serve()
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Println("run grpc serve failed, ", err)
		return
	}

	fmt.Println("run grpc serve success")
}
