package main

import (
	"log"
	"net/http"

	"github.com/genryusaishigikuni/micro-go/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"errors"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/genryusaishigikuni/micro-go/common/api"
	_ "github.com/joho/godotenv/autoload"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":8080")
	ordersServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.NewClient(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to orders service: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at ", ordersServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)
	log.Printf("Server listening on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start server")
	}

	srv := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-term
		if err := srv.Close(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error closing Server: %v", err)
		}
	}()

}
