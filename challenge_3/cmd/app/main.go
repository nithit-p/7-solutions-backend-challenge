package main

import (
	"flag"
	"log"

	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app"
	"github.com/nithit-p/7-solutions-backend-challenge/challenge_3/internal/app/beef"
)

var (
	host     = flag.String("host", "", "http server host")
	port     = flag.String("port", "8080", "http server port")
	grpcHost = flag.String("grpc_host", "localhost", "grpc server host")
	grpcPort = flag.String("grpc_port", "50051", "grpc server port")
)

func main() {
	flag.Parse()

	beefService, error := beef.NewBeefUsecase()
	if error != nil {
		log.Fatal("error while creating new BeefUsecase")
	}
	beefHttpHandler, err := beef.NewBeefHttpHandler(beefService)
	if err != nil {
		log.Fatal("error while creating new BeefHandler")
	}
	beefGrpcHandler, err := beef.NewBeefGrpcHandler(beefService)
	if err != nil {
		log.Fatal("error while creating new BeefGrpcHandler")
	}

	wait := make(chan struct{})
	// strarting http server
	go func() {
		httpAddr := *host + ":" + *port
		app.StartHttpServer(httpAddr, app.HttpHandlersConfig{
			BeefHandler: beefHttpHandler,
		})
	}()

	// starting grpc server
	go func() {
		grpcAddr := *grpcHost + ":" + *grpcPort
		app.StartGrpcServer(grpcAddr, app.GrpcHandlersConfig{
			BeefHandler: beefGrpcHandler,
		})
	}()
	<-wait
}
