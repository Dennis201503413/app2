package main

import (
	"context"
	"github.com/Dennis201503413/app2/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a := request.GetA()

	result := "Hellow World "+ a

	return &proto.Response{Result: result}, nil
}
