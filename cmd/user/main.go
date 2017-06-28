package main

import (
	_ "encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/dung13890/micro-go/pb/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	request []*pb.GetRequest
}

func (s *server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	// for _, user := range s.request {
	// 	if req.Keyword != "" {
	// 		fmt.Println(req.Keyword)
	// 	}
	// }
	users := &pb.GetResponse{
		Status: "success",
		Users: []*pb.Model{
			&pb.Model{
				Id:   "asdfsadfas",
				Name: req.Keyword,
			},
		},
	}
	return users, nil
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func main() {
	var port = flag.Int("port", 8080, "The server port")
	flag.Parse()
	fmt.Println(*port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterUserServer(srv, &server{})
	srv.Serve(lis)
}
