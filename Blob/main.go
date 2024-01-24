package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "block/protocol"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	name := req.GetName()
	message := "안녕하세요, " + name + "님! Go에서 보낸 메시지입니다."
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("리스닝 실패: %v", err)
	}

	srv := grpc.NewServer()

	log.Println("Go gRPC 서버가 50051 포트에서 실행 중입니다...")
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("서비스 실패: %v", err)
	}
}
