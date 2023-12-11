package main

import (
	"fmt"
	"github.com/rpuggian/ports-service/api/api"
	"github.com/rpuggian/ports-service/api/domain/grpc"
	"github.com/rpuggian/ports-service/api/domain/handlers"
	"github.com/rpuggian/ports-service/api/domain/services"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("starting API server")

	// Initialize grpc client connection
	conn, err := gogrpc.Dial(
		fmt.Sprintf(
			"%s:%s",
			os.Getenv("GRPC_SERVER_NAME"),
			os.Getenv("GRPC_PORT"),
		),
		// TODO: use secure credentials
		gogrpc.WithTransportCredentials(insecure.NewCredentials()),
		gogrpc.FailOnNonTempDialError(true),
		//block until connection is established
		gogrpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("could not get grpc connection %v", err)
	}
	defer conn.Close()

	// Initialize grpc client
	portClient := grpc.NewPortServiceClient(conn)

	// Initialize port service
	portService := services.NewPortService(portClient)

	// Initialize port handler
	portHandler := handlers.NewPortHandler(portService)

	// Initialize api server
	apiPort := os.Getenv("API_PORT")
	portServer := api.NewServer(apiPort, portHandler)

	// Serve api routes
	portServer.Serve()
	defer portServer.GracefulShutdown()
	log.Println("API server started, listening on port: ", apiPort)

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("shutting application down")

}
