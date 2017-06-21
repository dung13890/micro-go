package main

import (
	_ "encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	// pb "github.com/dung13890/micro-go/pb/user"
	_ "golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var port = flag.Int("port", 8080, "The server port")
	flag.Parse()
	fmt.Println(*port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	srv.Serve(lis)
}
