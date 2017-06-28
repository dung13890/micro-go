package main

import (
	_ "encoding/json"
	"fmt"
	"log"

	pb "github.com/dung13890/micro-go/pb/user"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// func getUsers(client pb.CustomerClient, customer *pb.CustomerRequest) {

// }

func main() {
	conn, err := grpc.Dial("user:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)
	request := &pb.GetRequest{
		Keyword: "testti su dan",
	}
	//fmt.Println(context.Background())

	resp, err := client.GetUsers(context.Background(), request)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(resp.Users[0].Name)
}
