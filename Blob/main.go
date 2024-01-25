package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "block/protocol"
)

type server struct {
    pb.UnimplementedBlockchainServiceServer
}

func (s *server) AddBlock(ctx context.Context, in *pb.BlockRequest) (*pb.BlockResponse, error) {
    // 블록 추가 로직 구현
    log.Printf("Received: %v", in.GetData())
    // 블록 추가 성공 시
    return &pb.BlockResponse{Status: "Success"}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterBlockchainServiceServer(s, &server{})
    log.Printf("Server listening at %v", lis.Addr())
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
