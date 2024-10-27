package services

import (
	"context"
	"time"

	"github.com/bohexists/product-list-svc/internal/proto"

	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	client proto.ProductServiceClient
}

func NewProductServiceClient(fetcherServiceAddr string) (*ProductServiceClient, error) {
	conn, err := grpc.Dial(fetcherServiceAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	client := proto.NewProductServiceClient(conn)
	return &ProductServiceClient{client: client}, nil
}

func (p *ProductServiceClient) FetchProducts(url string) (*proto.FetchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req := &proto.FetchRequest{Url: url}
	resp, err := p.client.Fetch(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
