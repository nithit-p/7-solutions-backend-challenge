package app

import (
	"log"
	"net"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/pb"

	"google.golang.org/grpc"
)

type GrpcHandlersConfig struct {
	BeefHandler pb.BeefServiceServer
}

func StartGrpcServer(addr string, handlersConfig GrpcHandlersConfig) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBeefServiceServer(grpcServer, handlersConfig.BeefHandler)

	// Start grpc server
	log.Printf("grpc server started at %s\n", addr)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}
