package main

import (
	"fmt"
	"os"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/cmd"
)

func main()  {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
}
