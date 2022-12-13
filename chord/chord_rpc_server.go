package chord

import (
	"context"
	"log"
)

type Server struct {
	CN ChordNode
}

func (s *Server) FindSuccessor(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	node := s.CN.findSuccessor(HashId{GetByteArrayFromString(in.Id)})
	return &NodeMessage{Log: "Hello From the Server!", Id: GetHexBasedStringFromBytes(node.node_id.Id), IpAddress: node.node_ip_address, Port: node.node_port_number}, nil
}

func (s *Server) FindPredecessor(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	node := s.CN.predecessor
	return &NodeMessage{Log: "Hello From the Server!", Id: GetHexBasedStringFromBytes(node.node_id.Id), IpAddress: node.node_ip_address, Port: node.node_port_number}, nil
}

func (s *Server) Notify(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)
	node := Node{node_id: HashId{GetByteArrayFromString(in.Id)}, node_ip_address: in.IpAddress, node_port_number: in.Port}
	return &NodeMessage{Log: "Hello From the Server!", Id: GetHexBasedStringFromBytes(node.node_id.Id), IpAddress: node.node_ip_address, Port: node.node_port_number, Success: true}, nil
}

func (s *Server) IsAlive(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	return &NodeMessage{Log: "Hello From the Server!", Success: true}, nil
}

func (s *Server) AddKeyValueToStore(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	success := s.CN.add_key_val_to_store(HashId{GetByteArrayFromString(in.Key)}, in.Value)
	return &NodeMessage{Log: "Hello From the Server!", Success: success}, nil
}

func (s *Server) IsKeyPresentInStore(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	found := s.CN.is_key_present_in_store(HashId{GetByteArrayFromString(in.Key)})
	return &NodeMessage{Log: "Hello From the Server!", Success: found}, nil
}

func (s *Server) GetValueOfKeyInStore(ctx context.Context, in *NodeMessage) (*NodeMessage, error) {
	log.Printf("Receive message body from client: %s %s %s %s %s", in.Id, in.IpAddress, in.Port, in.Path, in.Log)

	value, success := s.CN.get_value_of_key_in_store(HashId{GetByteArrayFromString(in.Key)})
	return &NodeMessage{Log: "Hello From the Server!", Value: value, Success: success}, nil
}

func (s *Server) mustEmbedUnimplementedChordServiceServer() {}
