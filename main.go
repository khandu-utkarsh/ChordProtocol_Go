package main

import (
	"fmt"
	"github.com/khandu-utkarsh/ChordProtocol_Go/chord"
)

func main() {

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
	fmt.Print("IP Adress: ", nodeFound.GetNodeIpAddress(), "\n")
	fmt.Print("Port Number: ", nodeFound.GetNodePortNumber(), "\n")

	return
}
