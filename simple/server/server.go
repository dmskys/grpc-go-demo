package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"grpcdemo/simple/pbproto"
	"golang.org/x/net/context"
)

// 定义服务端实现约定的接口
type UserInfoService struct{}
// 实现 interface
func (s *UserInfoService) GetUserInfo(content context.Context, req *pbproto.UserRequest) (resp *pbproto.UserResponse, err error) {
	uid := req.Uid
	// 模拟在数据库中查找用户信息
	if uid == 1 {
		resp = &pbproto.UserResponse{
			Id:    233,
			Name:  "wuYin",
			Age:   20,
			Title: []string{"Gopher", "PHPer"}, // repeated 字段是 slice 类型
		}
	}
	err = nil
	return
}

const (
	port = ":2333"
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

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}