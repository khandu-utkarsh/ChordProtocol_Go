package main

import (
	"fmt"
	"github.com/khandu-utkarsh/ChordProtocol_Go/chat"
	"github.com/khandu-utkarsh/ChordProtocol_Go/chord"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := chat.Server{}

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	fmt.Print("Utkarsh's chord-go code\n")

	currNodeIp := chord.GetCurrentProcessIPAddress()
	currNodePort := chord.GetCurrentProcessPort()

	stringToBeHashed := currNodeIp + currNodePort

	var currIpPortHash chord.HashId
	currIpPortHash = chord.Generate_Hash([]byte(stringToBeHashed))

	//!Initializing node data
	var currNode chord.Node
	currNode.SetNodeId(currIpPortHash)
	currNode.SetNodeIpAddress(currNodeIp)
	currNode.SetNodePortNumber(currNodePort)

	processNode := currNode.Create()

	var keyToFind chord.HashId
	nodeFound := processNode.Lookup(keyToFind)
	fmt.Print("IP Address: ", nodeFound.GetNodeIpAddress(), "\n")
	fmt.Print("Port Number: ", nodeFound.GetNodePortNumber(), "\n")

	return
}
