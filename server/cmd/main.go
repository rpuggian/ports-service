package main

import (
	"fmt"
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/grpc"
	"github.com/rpuggian/ports-service/server/domain/repository"
	"github.com/rpuggian/ports-service/server/domain/services"
	"github.com/rpuggian/ports-service/server/infra/redis"
	gogrpc "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting GRPC server")

	// Initialize redis grpc
	redisClient, err := redis.NewClient(os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"), "")
	if err != nil {
		log.Fatalf("failed to start redis grpc: %v", err)
	}

	// Initialize domain repository
	portRepository := repository.NewPortRepository(redisClient)

	// Initialize domain service
	portService := services.NewPortService(portRepository)

	// Initialize grpc server handler
	grpcServer := grpc.NewPortGRPCServer(portService)

	// Serve grpc
	startGRPC(grpcServer)

}

func startGRPC(server *grpc.PortGRPCServer) {
	grpcServer := gogrpc.NewServer()
	proto.RegisterPortServiceServer(grpcServer, server)

	//listen tcp domain
	grpcPort := os.Getenv("GRPC_PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//start grpc server
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("could not initialize GRPC server %v", err)
		}
	}()
	defer grpcServer.GracefulStop()
	log.Println("GRPC server started, listening on port: ", grpcPort)

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("shutting application down")
}
