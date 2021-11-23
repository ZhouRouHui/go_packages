package main

import (
	"context"
	"fmt"
	"go_packages/external/protobuf/pb"
	"google.golang.org/grpc"
)

func main() {
	/**
	这里是实现 grpc 客户端的代码
	*/

	// 连接 grpc 服务
	// grpc.WithInsecure() 表示以安全的方式进行操作
	grpcConn, err := grpc.Dial("127.0.0.1:9002", grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial grpc failed,", err)
		return
	}
	defer grpcConn.Close()

	// 初始化 grpc 客户端
	grpcClient := pb.NewHelloClient(grpcConn)

	// 调用远程服务。
	resp, err := grpcClient.HelloWorld(context.TODO(), &pb.Params{Name: "zrh"})
	if err != nil {
		fmt.Println("call hello world failed, err", err)
		return
	}

	fmt.Println(resp)
	fmt.Println("call hello world success")
}
