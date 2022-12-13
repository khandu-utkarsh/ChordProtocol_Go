package chord

import "fmt"

//import (
//	"fmt"
//)

func (n *Node) Create() ChordNode {

	var cn ChordNode
	cn.SelfNode = *n

	cn.predecessorStatus = false
	cn.predecessor = *n

	cn.successor = *n
	cn.fingerTable = InitializeFingerTable(*n)

	//!Initialize the store --Temp
	cn.store = InitializeStore()

	return cn
}

func (n *Node) Join(n_prime Node) ChordNode {

	var cn ChordNode
	cn.SelfNode = *n

	cn.predecessorStatus = false
	cn.predecessor = *n

	cn.successor = n_prime.RpcFindSuccessor(n.node_id)
	fmt.Println("SUCCESSOR: PORT: ", cn.successor.node_port_number)
	cn.fingerTable = InitializeFingerTable(cn.successor)

	//!Initialize the store
	cn.store = InitializeStore()

	return cn
}
