package main

import (
	"net"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"grpcdemo/simple/pbproto"
	"golang.org/x/net/context"
	grpc_transport "github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/kit/endpoint"
)

// 定义服务端实现约定的接口
type UserInfoService struct{
	userInfoHandler  grpc_transport.Handler
}

//通过grpc调用GetBookInfo时,GetBookInfo只做数据透传, 调用BookServer中对应Handler.ServeGRPC转交给go-kit处理
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pbproto.UserRequest) (resp *pbproto.UserResponse, err error) {

	_, rsp, err := s.userInfoHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	resp = rsp.(*pbproto.UserResponse)

	return
}

//创建userInfo的EndPoint
func makeGetUserInfoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//请求详情时返回 书籍信息
		req := request.(*pbproto.UserRequest)
		b := new(pbproto.UserResponse)
		b.Id = req.Uid
		b.Name = "21天精通php"
		b.Age = 21
		b.Title = []string{"title1", "title2"}
		return b,nil
	}
}

func decodeRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}


const (
	port = ":3333"
)

func main() {

	userInfoService := new(UserInfoService)
	//创建userInfo的Handler
	userInfoHandler := grpc_transport.NewServer(
		makeGetUserInfoEndpoint(),
		decodeRequest,
		encodeResponse,
	)
	//bookServer 增加 go-kit流程的 bookInfo处理逻辑
	userInfoService.userInfoHandler = userInfoHandler

	// 启动 gRPC 服务器。
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	fmt.Printf("listen %s\n", port)
	gs := grpc.NewServer()

	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	pbproto.RegisterUserInfoServiceServer(gs, userInfoService)

	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}