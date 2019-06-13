package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/api/v1"
	"context"
	"time"
	"github.com/golang/protobuf/ptypes"
)

const  (
	apiVersion = "v1"
)

func main()  {
	address := flag.String("server", "", "grpc server address host:port format")
	flag.Parse()

	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := v1.NewToDoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)

	remainder,_ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)


	req1 := v1.CreateRequest{
		Api: apiVersion,
		ToDo: &v1.ToDo{
			Title: "title ("+pfx+")",
			Description: "description",
			Reminder: remainder,
		},
	}

	res1, err := c.Create(ctx, &req1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res1)


	req2 := v1.ReadRequest{
		Api:apiVersion,
		Id: res1.Id,
	}
	res2, err := c.Read(ctx, &req2)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res2)

}
