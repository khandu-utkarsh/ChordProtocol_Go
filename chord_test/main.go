package main

import (
	"github.com/khandu-utkarsh/ChordProtocol_Go"
	"fmt"
)

func main() {

	fmt.Print("Utkarsh's chord-go code\n")

	curr_node_ip := GetCurrentProcessIPAddress()
	curr_node_port := GetCurrentProcessPort()

	stringToBeHashed := curr_node_ip + curr_node_port

	var curr_ip_port_hash HashId
	curr_ip_port_hash = Generate_Hash([]byte(stringToBeHashed))

	//!Initializing node data
	var curr_node Node

	curr_node.node_id = curr_ip_port_hash
	curr_node.node_ip_address = curr_node_ip
	curr_node.node_port_number = curr_node_port

	processNode := curr_node.Create()

	var keyToFind HashId
	nodeFound := processNode.lookup(keyToFind)
	fmt.Print("IP Adress: ", nodeFound.node_ip_address, "\n")
	fmt.Print("Port Number: ", nodeFound.node_port_number, "\n")

	return
}
