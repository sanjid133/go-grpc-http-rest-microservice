package grpc

import (
	"context"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/api/v1"
	"net"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"log"
)

func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, port string) error  {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	v1.RegisterToDoServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	log.Println("starting grpc server ...")
	return server.Serve(listen)
}