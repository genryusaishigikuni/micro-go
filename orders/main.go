package main

import (
	"context"
	"log"
	"net"

	"github.com/genryusaishigikuni/micro-go/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())

	log.Printf("GRPC Server Started at %s", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf(err.Error())
	}

}
