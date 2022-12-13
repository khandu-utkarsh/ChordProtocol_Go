package main

import (
	"flag"
	"fmt"
	"github.com/khandu-utkarsh/ChordProtocol_Go/chord"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func checkIPAddress(ip string) {
	if ip != "" && ip != "localhost" {
		if net.ParseIP(ip) == nil {
			log.Fatalf("IP Address: %s - Invalid\n", ip)
		} else {
			fmt.Printf("IP Address: %s - Valid\n", ip)
		}
	}
}

func main() {
	// accept IP address and port
	// check if creating a new network or joining existing network with ipAddress and port of existing node
	ipAddress := flag.String("ip_address", "localhost", "IP Address")
	port := flag.String("port", "9000", "Port")

	joinIpAddress := flag.String("nw_ip_address", "", "Network IP Address")
	joinPort := flag.String("nw_port", "", "Network IP Address")
	flag.Parse()

	fmt.Println("Current IP Address:port - ", *ipAddress, ":", *port)
	fmt.Println("Network IP Address:port - ", *ipAddress, ":", *joinPort)

	checkIPAddress(*ipAddress)
	checkIPAddress(*joinIpAddress)

	fmt.Print("Utkarsh's chord-go code\n")

	currNodeIp := *ipAddress
	currNodePort := *port

	stringToBeHashed := currNodeIp + currNodePort

	var currIpPortHash chord.HashId
	currIpPortHash = chord.Generate_Hash([]byte(stringToBeHashed))

	//!Initializing node data
	var currNode chord.Node
	currNode.SetNodeId(currIpPortHash)
	currNode.SetNodeIpAddress(currNodeIp)
	currNode.SetNodePortNumber(currNodePort)

	var (
		processNode chord.ChordNode
	)

	if *joinPort != "" && *joinIpAddress != "" {

		joinNode := chord.Node{}
		joinNode.SetNodeId(chord.Generate_Hash([]byte(*joinIpAddress + *joinPort)))
		joinNode.SetNodeIpAddress(*joinIpAddress)
		joinNode.SetNodePortNumber(*joinPort)
		processNode = currNode.Join(joinNode)
	} else {
		processNode = currNode.Create()
	}

	fmt.Println("IP/Port")
	chord.PrintBytesSplices(processNode.SelfNode.GetNodeId().Id)

	// Start server for RPCs
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *ipAddress, *port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		s := chord.Server{CN: processNode}

		grpcServer := grpc.NewServer()
		chord.RegisterChordServiceServer(grpcServer, &s)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}()

	// Stabilise
	go func() {
		processNode.PerodicallyCheck()
	}()

	// Open for consumer requests...
	go func() {
		time.Sleep(10 * time.Second)
		processNode.PrintFingerTable()
	}()

	var keyToFind chord.HashId
	keyToFind = chord.Generate_Hash([]byte("DistributedSystems"))
	fmt.Println("Key: ")
	chord.PrintBytesSplices(keyToFind.Id)
	nodeFound := processNode.Lookup(keyToFind)
	fmt.Print("IP Address: ", nodeFound.GetNodeIpAddress(), "\n")
	fmt.Print("Port Number: ", nodeFound.GetNodePortNumber(), "\n")
	select {}
	return
}
