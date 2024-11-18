package main

import (
	"context"
	"log"

	pb "github.com/genryusaishigikuni/micro-go/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received!")
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
