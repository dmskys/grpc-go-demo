package main

import (
	"google.golang.org/grpc"
	"log"
	"grpcdemo/serversidestream/pbproto"
	"golang.org/x/net/context"
	"io"
)

const (
	port = ":5333"
)
func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v\n", err)
	}
	defer conn.Close()

	// 实例化 UserInfoService 微服务的客户端
	client := pbproto.NewUserInfoServiceClient(conn)

	// 调用服务
	req := new(pbproto.UserRequest)
	req.Uid = 1

	stream, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalf("recevie resp error: %v", err)
	}

	// 接收流数据
	for {
		resp, err := stream.Recv()
		log.Printf("err : %v", err)
		if err == io.EOF { // 服务端数据发送完毕
			break
		}
		if err != nil {
			log.Fatalf("receive error: %v", err)
		}
		log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应
	}

}