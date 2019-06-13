package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/api/v1"
	"net/http"
	"os/signal"
	"os"
	"time"
	"log"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error  {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		return err
	}

	srv := &http.Server{
		Addr: ":"+httpPort,
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {

		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()

	log.Println("starting http/rest srever...")
	return srv.ListenAndServe()
}
