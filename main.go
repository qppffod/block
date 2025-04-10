package main

import (
	"fmt"
	"log"
	"net"

	"github.com/qppfod/block/node"
	"github.com/qppfod/block/proto"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	node := node.NewNode()

	proto.RegisterNodeServer(grpcServer, node)
	fmt.Println("node running on port :3000")
	grpcServer.Serve(ln)
}
