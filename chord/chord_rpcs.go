package chord

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func (sendNode *Node) getNewRpcConnection() *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", sendNode.node_ip_address, sendNode.node_port_number), grpc.WithInsecure())
	if err != nil {
		//log.Fatalf("did not connect: %s", err)

		// TODO : Do things based on failure of connection
	}
	return conn
}

func (sendNode *Node) RpcFindSuccessor(id HashId) Node {

	//node should call find_successor funcation on it's machine and get this value
	//var nodeReturned Node
	//return nodeReturned
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.FindSuccessor(context.Background(), &NodeMessage{Id: GetHexBasedStringFromBytes(id.Id)})
	if err != nil {
		log.Fatalf("Error when calling FindSuccessor: %s", err)
	}
	log.Printf("RpcFindSuccessor : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	node := Node{node_id: HashId{GetByteArrayFromString(response.Id)}, node_ip_address: response.IpAddress, node_port_number: response.Port}

	return node
}

func (sendNode *Node) RpcFindPredecessor() Node {
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.FindPredecessor(context.Background(), &NodeMessage{})
	if err != nil {
		log.Fatalf("Error when calling FindPredecessor: %s", err)
	}
	log.Printf("RpcFindPredecessor : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	node := Node{node_id: HashId{GetByteArrayFromString(response.Id)}, node_ip_address: response.IpAddress, node_port_number: response.Port}

	return node
}

func (sendNode *Node) RpcNotify(receive_node Node) bool {

	//!node should call notify on it with n_prime as argument, I'm the new successor
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.Notify(context.Background(), &NodeMessage{Id: GetHexBasedStringFromBytes(receive_node.node_id.Id), IpAddress: receive_node.node_ip_address, Port: receive_node.node_port_number})
	if err != nil {
		log.Fatalf("Error when calling Notify: %s", err)
	}
	log.Printf("RpcNotify : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	return response.Success
}

func (sendNode *Node) RpcIsAlive(receive_node Node) bool {
	//! send_node should send messages to receive node to check if it is alive or not,
	// if yes, return ture and if not return false
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.Notify(context.Background(), &NodeMessage{})
	if err != nil {
		return false
	}
	log.Printf("RpcIsAlive : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	return response.Success
}

func (sendNode *Node) RpcAddKeyValueToStore(key HashId, value string) bool {
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.AddKeyValueToStore(context.Background(), &NodeMessage{Key: GetHexBasedStringFromBytes(key.Id), Value: value})
	if err != nil {
		log.Fatalf("Error when calling AddKeyValueToStore: %s", err)
	}
	log.Printf("RpcAddKeyValueToStore : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	return response.Success
}

func (sendNode *Node) RpcIsKeyPresentInStore(key HashId) bool {
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.IsKeyPresentInStore(context.Background(), &NodeMessage{Key: GetHexBasedStringFromBytes(key.Id)})
	if err != nil {
		log.Fatalf("Error when calling IsKeyPresentInStore: %s", err)
	}
	log.Printf("RpcIsKeyPresentInStore : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	return response.Success
}

func (sendNode *Node) RpcGetValueOfKeyInStore(key HashId) (string, bool) {
	conn := sendNode.getNewRpcConnection()
	defer conn.Close()

	c := NewChordServiceClient(conn)
	response, err := c.GetValueOfKeyInStore(context.Background(), &NodeMessage{Key: GetHexBasedStringFromBytes(key.Id)})
	if err != nil {
		log.Fatalf("Error when calling GetValueOfKeyInStore: %s", err)
	}
	log.Printf("RpcGetValueOfKeyInStore : Response from server: %s %s %s %s %s", response.Id, response.Port, response.IpAddress, response.Path, response.Log)

	return response.Value, response.Success
}
