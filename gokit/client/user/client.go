package main

import (
	"google.golang.org/grpc"
	"fmt"
	"log"
	"grpcdemo/simple/pbproto"
	"golang.org/x/net/context"
)

func main() {

	// 不使用认证建立连接
	conn, err := grpc.Dial(":3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	// 实例化 UserInfoService 微服务的客户端  创建 gRPC 客户端实例
	client := pbproto.NewUserInfoServiceClient(conn)

	// 调用服务
	req := new(pbproto.UserRequest)
	req.Uid = 1

	resp, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("resp error: %v\n", err)
	}

	fmt.Printf("Recevied: %v\n", resp)
}