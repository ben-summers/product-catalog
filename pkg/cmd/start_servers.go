package cmd

import (
	"context"
	"log"
	"product-catalog/pkg/api/service"
	"product-catalog/pkg/config"
	"product-catalog/pkg/server/grpc"
	"product-catalog/pkg/server/rest"
)

func StartServers() error {

	if err := config.Configure(); err != nil {
		log.Fatalf("%+v", err)
	}

	ctx := context.Background()
	svc := service.NewApiService()

	go func() {
		if err := rest.StartServer(ctx, "50051", "8080"); err != nil {
			log.Panic(err)
		}

	}()

	if err := grpc.StartServer(ctx, svc, "50051"); err != nil {
		return err
	}

	return nil
}
