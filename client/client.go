package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/khandu-utkarsh/ChordProtocol_Go/chord"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("192.168.1.94:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chord.NewChordServiceClient(conn)

	response, err := c.FindSuccessor(context.Background(), &chord.NodeMessage{Log: "Testing", Id: "#343434", IpAddress: "192.168.59.95"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Log)
}
