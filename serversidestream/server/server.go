package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"google.golang.org/grpc/reflection"
	"grpcdemo/serversidestream/pbproto"
)

// 模拟的数据库查询结果
var users = map[int32]pbproto.UserResponse{
	1: {Name: "Dennis MacAlistair Ritchie", Age: 70},
	2: {Name: "Ken Thompson", Age: 75},
	3: {Name: "Rob Pike", Age: 62},
}

// 定义服务端实现约定的接口
type UserInfoService struct{}
// 实现 interface
func (s *UserInfoService) GetUserInfo(req *pbproto.UserRequest, stream pbproto.UserInfoService_GetUserInfoServer)  error {
	//uid := req.Uid
	// 响应流数据
	for _, user := range users {
		stream.Send(&user)
	}
	log.Printf("[RECEIVED REQUEST]: %v\n", req)
	return nil
}

const (
	port = ":5333"
)

func main() {
	// 启动 gRPC 服务器。
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	s := grpc.NewServer()

	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	pbproto.RegisterUserInfoServiceServer(s, &UserInfoService{})

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}