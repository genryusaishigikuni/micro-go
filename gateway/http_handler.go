package main

import (
	"net/http"

	"github.com/genryusaishigikuni/micro-go/common"
	pb "github.com/genryusaishigikuni/micro-go/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customer_id}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
