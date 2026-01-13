package main

import (
	"log"

	"github.com/Mellanie-Marques/microservices/order/config"
	"github.com/Mellanie-Marques/microservices/order/internal/adapter/db"
	"github.com/Mellanie-Marques/microservices/order/internal/adapter/grpc"
	"github.com/Mellanie-Marques/microservices/order/internal/adapters/payment"
	"github.com/Mellanie-Marques/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
