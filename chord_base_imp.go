package chord

import (
	"fmt"
)

// Constants
const m = 6

//!Code has to be something like
// Every process will be a node and will have a finger table

//SHA 1 Algorithm is only returning the const BLOCK_SIZE hash, check if there is a way to change the BLOCK_SIZE to m

type FingerTableEntry struct {
	node_ip_address  string
	node_port_number string
}

// Finger type denoting identifying information about a ChordNode
type FingerTable struct {
	table [m]FingerTableEntry //m entities
}

type Node struct {
	node_hash_id [m]byte
	successor * Node
	predecessor * Node
}

// Implement the comparison function
func IsIdBetweenRange_RightEnd_Inclusive(key []byte, min []byte, max []byte) bool {
}

func IsIdBetweenRange_RightEnd_Exclusive(key []byte, min []byte, max []byte) bool {

}

func closest_preceding_node(key_hash_id []byte) []byte []
	


func (node Node) closest_preceding_node(key_hash_id []byte) Node {
	for i := m; i >= 1; i-- {

		node_at_index_i_id := node.fingerTable.table[i].associated_node_hash_id
		if node_at_index_i_id > node.node_hash_id && node_at_index_i_id < key_hash_id {
			return *node.associated_node
		}
	}
	return node
}



// Functions from Image 5 from the paper read in Notability app
func (node Node) find_successor(key_hash_id []byte) Node {
	if IsIdBetweenRange(key_hash_id, node.node_hash_id, node.successor.node_hash_id) {
		return *node.successor
	} else {
		n_prime = node.closest_preceding_node(key_hash_id)
		//!Now a message should be sent to n_prime node to invoke find_successor operation
		//!Therefore below function would be bringing results from different process  
		return n_prime.find_successor(key_hash_id) //This function should be basically be searching for other machine node
	}
}


// Functions from the Image 6 from the paper read in Notability app
func (node Node) create() {
	node.predecessor = nil
	node.successor = &node
}

func (node Node) join(node_prime Node) {
	node.predecessor = nil
	returnedNode := node_prime.find_successor(node.node_hash_id)
	node.successor = &returnedNode
}


func (node Node) stablize() {
	x := node.successor.predecessor
	if x
}

func main() {
	fmt.Println("Hello, world.")
}
