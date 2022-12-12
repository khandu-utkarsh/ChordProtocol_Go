package main

import (
	"fmt"
	"github.com/khandu-utkarsh/ChordProtocol_Go/chord"
)

func main() {

	fmt.Print("Utkarsh's chord-go code\n")

	curr_node_ip := chord.GetCurrentProcessIPAddress()
	curr_node_port := chord.GetCurrentProcessPort()

	stringToBeHashed := curr_node_ip + curr_node_port

	var curr_ip_port_hash chord.HashId
	curr_ip_port_hash = chord.Generate_Hash([]byte(stringToBeHashed))

	//!Initializing node data
	var curr_node chord.Node
	curr_node.Set_node_id(curr_ip_port_hash)
	curr_node.Set_node_ip_address(curr_node_ip)
	curr_node.Set_node_port_number(curr_node_port)

	processNode := curr_node.Create()

	var keyToFind chord.HashId
	nodeFound := processNode.Lookup(keyToFind)
	fmt.Print("IP Adress: ", nodeFound.Get_node_ip_address(), "\n")
	fmt.Print("Port Number: ", nodeFound.Get_node_port_number(), "\n")

	return
}
