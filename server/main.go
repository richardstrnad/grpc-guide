package main

import (
	"context"
	"log"
	"net"

	pb "grpc-gcloud/ping"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPingerServer
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
	log.Printf("Received: %v", in.Message)
	return &pb.PingReply{Message: "Hello " + in.Message}, nil
}

func (s *server) GetVersion(ctx context.Context, in *pb.VersionRequest) (*pb.VersionReply, error) {
	return &pb.VersionReply{Message: "0.0.1"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPingerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
