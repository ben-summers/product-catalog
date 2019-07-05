package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"product-catalog/gen/api"
)

func StartServer(ctx context.Context, service api.ApiServiceServer, port string) error {
	if listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port)); err != nil {
		return err
	} else {
		server := grpc.NewServer()
		api.RegisterApiServiceServer(server, service)

		// graceful shutdown
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				log.Println("shutting down grpc server...")
				server.GracefulStop()

				<-ctx.Done()
			}
		}()

		// start server
		log.Println("starting grpc...")
		return server.Serve(listen)

	}
}
