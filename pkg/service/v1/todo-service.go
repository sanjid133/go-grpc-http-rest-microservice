package v1

import (
	"database/sql"
	"github.com/sanjid133/go-grpc-http-rest-microservice/pkg/api/v1"
	"context"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"github.com/golang/protobuf/ptypes"
	"fmt"
	"time"
)

const (
	apiVersion = "v1"
)

type toDoServiceServer struct {
	db *sql.DB
}

func NewToDoServiceServer(db *sql.DB) v1.ToDoServiceServer  {
	return &toDoServiceServer{db}
}

func (s *toDoServiceServer) checkApi(api string) error  {
	if len(api) > 0 {
		if api != apiVersion {
			return status.Errorf(codes.Unimplemented,
				"unsupported api version %v", api)
		}
	}
	// api="" means current version
	return nil
}

func (s *toDoServiceServer) connect(ctx context.Context) (*sql.Conn, error)  {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Error(codes.Unknown, "filed to connect")
	}
	return c, nil
}

func (s *toDoServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error)  {
	if err := s.checkApi(req.Api); err != nil {
		return nil, err
	}
	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	reminder, err := ptypes.Timestamp(req.ToDo.Reminder)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "remainer is not valid")
	}

	res, err := conn.ExecContext(ctx,
		"INSERT INTO ToDo(`Title`, `Description`, `Remainder`) values(?, ?, ?)",
			req.ToDo.Title, req.ToDo.Description, reminder)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &v1.CreateResponse{
		Api:apiVersion,
		Id:id,
	}, nil
}

func (s *toDoServiceServer) Read(ctx context.Context, req *v1.ReadRequest) (*v1.ReadResponse, error)  {
	if err := s.checkApi(req.Api); err != nil {
		return nil, err
	}
	conn, err := s.connect(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()


	rows, err := conn.QueryContext(ctx,
		"SELECT `ID`, `Title`, `Description`, `Remainder` from ToDo where `ID`=?",
		req.Id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, err
		}
		return nil, status.Error(codes.NotFound, fmt.Sprintf("no rows found with id %v", req.Id))
	}

	var td v1.ToDo
	var remainder time.Time

	if err := rows.Scan(&td.Id, &td.Title, &td.Description, &remainder); err != nil {
		return nil, err
	}
	td.Reminder, err = ptypes.TimestampProto(remainder)

	if rows.Next() {
		return nil, status.Error(codes.Unknown, fmt.Sprintf("found multiple rows with id %v", req.Id))
	}

	return &v1.ReadResponse{
		Api:apiVersion,
		ToDo: &td,
	}, nil


}

func (s *toDoServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error)  {
	return nil, nil
}

func (s *toDoServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error)  {
	return nil, nil
}

func (s *toDoServiceServer) ReadAll(ctx context.Context, req *v1.ReadAllRequest) (*v1.ReadAllResponse, error)  {
	return nil, nil
}