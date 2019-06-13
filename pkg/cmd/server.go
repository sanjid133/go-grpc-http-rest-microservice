package cmd

import (
	"context"
	"flag"
	"fmt"
	"database/sql"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/service/v1"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/protocol/grpc"
)

type Database struct {
	Host string
	User string
	Password string
	Scheme string
}

type Config struct {
	Port string
	DB Database
}

func RunServer() error  {
	ctx := context.Background()

	var cfg Config
	cfg.DB = Database{}
	flag.StringVar(&cfg.Port, "grpc-port", "8080", "grpc port to bind")
	flag.StringVar(&cfg.DB.Host, "host", "localhost", "database host")
	flag.StringVar(&cfg.DB.User, "user", "", "database user")
	flag.StringVar(&cfg.DB.Password, "password", "", "database password")
	flag.StringVar(&cfg.DB.Scheme, "scheme", "", "database scheme")

	flag.Parse()

	param := "parseTime=true"

	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Scheme,
		param,
		)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}
	defer db.Close()

	v1api := v1.NewToDoServiceServer(db)
	return grpc.RunServer(ctx, v1api, cfg.Port)
}


