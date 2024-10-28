package main

import (
	"fmt"
	"github.com/bohexists/product-list-svc/internal/config"
	"github.com/bohexists/product-list-svc/internal/services"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	productClient, err := services.NewProductServiceClient(cfg.FetcherServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to product-fetcher-svc: %v", err)
	}

	url := "http://example.com/products.csv"

	resp, err := productClient.FetchProducts(url)
	if err != nil {
		log.Fatalf("Failed to fetch products: %v", err)
	}

	fmt.Println(resp)
}
