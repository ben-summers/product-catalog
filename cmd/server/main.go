package main

import (
	"context"
	"log"
	"product-catalog/pkg/api/service"
	"product-catalog/pkg/server/grpc"
	"product-catalog/pkg/server/rest"
)

func main() {

	ctx := context.Background()
	svc := service.NewApiService()

	go func() {
		if err := rest.StartServer(ctx, "50051", "8080"); err != nil {
			log.Panic(err)
		}

	}()

	if err := grpc.StartServer(ctx, svc, "50051"); err != nil {
		log.Panic(err)
	}

}
