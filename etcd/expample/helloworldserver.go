package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpclb "grpcdemo/etcd/balance"
	"grpcdemo/etcd/expample/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	servs = flag.String("service", "hello_service", "service name")
	port = flag.String("port", "50001", "listening port")
	regs  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	defer lis.Close()

	gs := grpc.NewServer()
	defer gs.GracefulStop()

	pb.RegisterGreeterServer(gs, &server{})


	err = grpclb.Register(*servs, "127.0.0.1", *port, *regs, time.Second*10, 15)
	if err != nil {
		log.Fatalf("grpclb.Register: %s", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		grpclb.UnRegister()
		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}

	}()

	log.Printf("starting hello service at %s", *port)

	if err := gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}