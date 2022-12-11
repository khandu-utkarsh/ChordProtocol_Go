package chord

//import (
//	"fmt"
//)

func (n Node) Create() ChordNode {

	var cn ChordNode
	cn.self_node = n

	cn.predecessorStatus = false
	cn.predecessor = n

	cn.successor = n
	cn.fingerTable = InitializeFingerTable(n)
	return cn
}

func (n Node) Join(n_prime Node) ChordNode {

	var cn ChordNode
	cn.self_node = n

	cn.predecessorStatus = false
	cn.predecessor = n

	cn.successor = n_prime.RPC_find_successor(n.node_id)
	cn.fingerTable = InitializeFingerTable(cn.successor)
	return cn
}
