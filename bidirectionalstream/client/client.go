package main

import (
	"google.golang.org/grpc"
	"log"
	"grpcdemo/bidirectionalstream/pbproto"
	"golang.org/x/net/context"
	"io"
)

const (
	port = ":2335"
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
	// 模拟的数据库中有 3 条记录，ID 分别为 1 2 3 发送流数据
	for i = 1; i < 4; i++ {
		err := stream.Send(&pbproto.UserRequest{Uid: i})
		if err != nil {
			log.Fatalf("send error: %v", err)
		}
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
	stream.CloseSend()
}