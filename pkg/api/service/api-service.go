package service

import (
	"context"
	"product-catalog/gen/api"
)

type apiService struct {
}

func (apiService) GetProduct(ctx context.Context, req *api.GetProductRequest) (*api.GetProductResponse, error) {
	return &api.GetProductResponse{
		ID:      req.ID,
		Product: nil,
	}, nil
}

func NewApiService() api.ApiServiceServer {
	return &apiService{}
}
