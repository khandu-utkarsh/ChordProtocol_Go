package chord

import (
	"crypto/sha1"
	"time"
)

// Constants
// !We need m bit identifier, SHA1 returns 20 * 8 bits = 160 bits
const m = int(sha1.Size * 8)

// type Key struct {
// 	key_id HashId
// }

// ! Each process on the system will be identified by type Node -  abstraction over Id and port number and hashId for that process
type Node struct {
	node_id          HashId
	node_ip_address  string
	node_port_number string
}

func (node Node) Set_node_id(id HashId) {
	node.node_id = id
}

func (node Node) Set_node_ip_address(ip string) {
	node.node_ip_address = ip
}

func (node Node) Set_node_port_number(p string) {
	node.node_port_number = p
}

func (node Node) Get_node_ip_address() (string) {
	return node.node_ip_address
}

func (node Node) Get_node_port_number() (string) {
	return node.node_port_number
}


// Finger type denoting identifying information about a ChordNode
type FingerTable struct {
	table [m]Node //m entities =  159 entries and successor will directly be stored
	next  int
}

//!Every machine (processor) will be of type ChordNode
//!It can be imagined as on the system each node will identity as ChordNode
//!So when a processor joins the node ring or even creates a new ring, a ChordNode object for that will be returned

// !Cord Node will be basically current node (current process)
type ChordNode struct {
	fingerTable FingerTable
	self_node   Node

	successor Node //!Should always be pointer to the first element of the finger table

	//!If status is false, mean predecessor is some garbage value so consider it nil according to paper
	predecessorStatus bool //If status is false,
	predecessor       Node
}

func (cn ChordNode) UpdateSuccessor(node Node) {
	cn.successor = node
	cn.fingerTable.table[0] = node
}

func (ch ChordNode) Lookup(key HashId) Node {
	successorFound := ch.find_successor(key)
	return successorFound
}

func InitializeFingerTable(node Node) FingerTable {
	var ftable FingerTable
	ftable.next = 1

	for i := 0; i < m; i++ {
		ftable.table[i] = node
	}
	return ftable
}

func (cn ChordNode) closest_preceding_node(key HashId) Node {
	for i := m - 1; i >= 0; i-- {
		if IsIdBetweenRange_RightEnd_Exclusive(cn.fingerTable.table[i].node_id, cn.self_node.node_id, key) {
			return cn.fingerTable.table[i]
		}
	}
	return cn.self_node
}

func (cn ChordNode) find_successor(key HashId) Node {
	if IsIdBetweenRange_RightEnd_Inclusive(key, cn.self_node.node_id, cn.successor.node_id) {
		return cn.successor
	}
	//!This is the else condition, since if above if is true, this won't be executed
	n_prime := cn.closest_preceding_node(key)
	return n_prime.RPC_find_successor(key)
}

//!Following three functions are called periodically
//	1. stabilize
//	2. fix_fingers
//	3. check_predecessor

func (cn ChordNode) stabilize() {

	predecessorOfSuccessor := cn.successor.RPC_find_predecessor()

	if IsIdBetweenRange_RightEnd_Exclusive(predecessorOfSuccessor.node_id, cn.self_node.node_id, cn.successor.node_id) {
		cn.UpdateSuccessor(predecessorOfSuccessor)
	}
	//!Notifying it's successor about it
	cn.successor.RPC_notify(cn.self_node) //!This will be a RPC Call
}

func (ch ChordNode) notify(n_prime Node) {
	if ch.predecessorStatus == false || IsIdBetweenRange_RightEnd_Exclusive(n_prime.node_id, ch.predecessor.node_id, ch.self_node.node_id) == true {
		ch.predecessor = n_prime
	}
}

func (ch ChordNode) check_predecessor() {
	isAlive := ch.self_node.RPC_IsAlive(ch.predecessor)
	if isAlive == false { //!Mean node has failed
		ch.predecessorStatus = false
	}
}

func (cn ChordNode) fix_fingers() {

	cn.fingerTable.next = cn.fingerTable.next + 1

	if cn.fingerTable.next >= m {
		cn.fingerTable.next = 0
	}

	//!Generate hash id for the number
	var int_hash_id HashId //!Write a message to fetch this id
	int_hash_id = GenerateHashIdForFingerIndex(cn.self_node.node_id, cn.fingerTable.next)

	newSuccessorReturned := cn.find_successor(int_hash_id)
	cn.fingerTable.table[cn.fingerTable.next] = newSuccessorReturned
}

//!Write a function which checks what timer has went off and then do as instructed

// !This will be infinitely running
func (ch ChordNode) perodicallyCheck() {
	//!Create three timers for stablization, fix fingers and check_predecessor
	stable_ticker := time.NewTicker(5 * time.Millisecond)
	check_p_ticker := time.NewTicker(4 * time.Second)
	fix_f_timer := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-stable_ticker.C:
			ch.stabilize()

		case <-check_p_ticker.C:
			ch.check_predecessor()

		case <-fix_f_timer.C:
			ch.fix_fingers()
		}
	}
}
