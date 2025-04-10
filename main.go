package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/qppfod/block/node"
	"github.com/qppfod/block/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ln, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	node := node.NewNode()

	proto.RegisterNodeServer(grpcServer, node)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			makeTransaction()
		}
	}()

	fmt.Println("node running on port :4000")
	grpcServer.Serve(ln)
}

func makeTransaction() {
	client, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := proto.NewNodeClient(client)

	version := &proto.Version{
		Version: "blocker-0.1",
		Height:  1,
	}

	_, err = c.Handshake(context.TODO(), version)
	if err != nil {
		log.Fatal(err)
	}
}
