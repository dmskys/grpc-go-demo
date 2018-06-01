package main

import (
	"google.golang.org/grpc"
	"log"
	"grpcdemo/clientsidestream/pbproto"
	"golang.org/x/net/context"
)

const (
	port = ":2334"
)
func main() {

	// 不使用认证建立连接
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

	// 向服务端发送流数据
	stream, err := client.GetUserInfo(context.Background())

	var i int32
	// 模拟的数据库中有 3 条记录，ID 分别为 1 2 3
	for i = 1; i < 4; i++ {
		err := stream.Send(&pbproto.UserRequest{Uid: i})
		if err != nil {
			log.Fatalf("send error: %v", err)
		}
	}

	// 接收服务端的响应
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("recevie resp error: %v", err)
	}

	log.Printf("[RECEIVED RESPONSE]: %v\n", resp) // 输出响应



}