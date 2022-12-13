package chord

import (
	"crypto/sha1"
	"time"
)

// Constants
// !We need m bit identifier, SHA1 returns 20 * 8 bits = 160 bits
const spliceElementsCount = int(sha1.Size)
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

func (node *Node) SetNodeId(id HashId) {
	node.node_id = id
}

func (node *Node) SetNodeIpAddress(ip string) {
	node.node_ip_address = ip
}

func (node *Node) SetNodePortNumber(p string) {
	node.node_port_number = p
}

func (node *Node) GetNodeId() HashId {
	return node.node_id
}

func (node *Node) GetNodeIpAddress() string {
	return node.node_ip_address
}

func (node *Node) GetNodePortNumber() string {
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
	SelfNode    Node

	successor Node //!Should always be pointer to the first element of the finger table

	//!If status is false, mean predecessor is some garbage value so consider it nil according to paper
	predecessorStatus bool //If status is false,
	predecessor       Node

	//!Adding the key value pair data structure to store the keys
	//!Since splices can be used, hence convert them to big.Int for storage
	store map[string]string
}

func (cn *ChordNode) UpdateSuccessor(node Node) {
	cn.successor = node
	cn.fingerTable.table[0] = node
}

func (cn *ChordNode) Query_AddKeyValueToStore(key HashId, value string) bool {
	node := cn.Lookup(key) //!Successor found
	if GetHexBasedStringFromBytes(node.GetNodeId().Id) == GetHexBasedStringFromBytes(cn.SelfNode.GetNodeId().Id) {
		return cn.add_key_val_to_store(key, value)
	}

	additionStatus := node.RpcAddKeyValueToStore(key, value)
	return additionStatus
}

func (cn *ChordNode) Query_IsKeyPresentInStore(key HashId) bool {
	node := cn.Lookup(key) //!Successor found
	if GetHexBasedStringFromBytes(node.GetNodeId().Id) == GetHexBasedStringFromBytes(cn.SelfNode.GetNodeId().Id) {
		return cn.is_key_present_in_store(key)
	}
	presentStatus := node.RpcIsKeyPresentInStore(key)
	return presentStatus
}

func (cn *ChordNode) Query_GetValueOfKeyInStore(key HashId) (string, bool) {
	node := cn.Lookup(key) //!Successor found
	if GetHexBasedStringFromBytes(node.GetNodeId().Id) == GetHexBasedStringFromBytes(cn.SelfNode.GetNodeId().Id) {
		return cn.get_value_of_key_in_store(key)
	}
	value, status := node.RpcGetValueOfKeyInStore(key)
	return value, status
}

func (cn *ChordNode) add_key_val_to_store(key HashId, value string) bool {
	hashStr := GetHexBasedStringFromBytes(key.Id)
	cn.store[hashStr] = value
	return true
}

func (cn *ChordNode) is_key_present_in_store(key HashId) bool {
	hashStr := GetHexBasedStringFromBytes(key.Id)
	_, ok := cn.store[hashStr]
	return ok
}

func (cn *ChordNode) get_value_of_key_in_store(key HashId) (string, bool) {
	hashStr := GetHexBasedStringFromBytes(key.Id)
	out, ok := cn.store[hashStr]
	return out, ok
}

func (ch *ChordNode) Lookup(key HashId) Node {
	successorFound := ch.findSuccessor(key)
	return successorFound
}

// !Write function to store key and return the ip address, just to check if it storing it correctly
func InitializeStore() map[string]string {
	store := make(map[string]string)
	return store
}

func InitializeFingerTable(node Node) FingerTable {
	var fTable FingerTable
	fTable.next = 1

	for i := 0; i < m; i++ {
		fTable.table[i] = node
	}
	return fTable
}

func (cn *ChordNode) closestPrecedingNode(key HashId) Node {
	for i := m - 1; i >= 0; i-- {
		if IsIdBetweenRangeRightEndExclusive(cn.fingerTable.table[i].node_id, cn.SelfNode.node_id, key) {
			return cn.fingerTable.table[i]
		}
	}
	return cn.SelfNode
}

func (cn *ChordNode) findSuccessor(key HashId) Node {
	if IsIdBetweenRange_RightEnd_Inclusive(key, cn.SelfNode.node_id, cn.successor.node_id) {
		return cn.successor
	}
	//!This is the else condition, since if above if is true, this won't be executed
	n_prime := cn.closestPrecedingNode(key)
	if GetHexBasedStringFromBytes(n_prime.GetNodeId().Id) == GetHexBasedStringFromBytes(cn.SelfNode.GetNodeId().Id) {
		return cn.SelfNode
	}

	return n_prime.RpcFindSuccessor(key)
}

//!Following three functions are called periodically
//	1. stabilize
//	2. fix_fingers
//	3. check_predecessor

func (cn *ChordNode) stabilize() {

	predecessorOfSuccessor := cn.successor.RpcFindPredecessor()

	if IsIdBetweenRangeRightEndExclusive(predecessorOfSuccessor.node_id, cn.SelfNode.node_id, cn.successor.node_id) {
		cn.UpdateSuccessor(predecessorOfSuccessor)
	}
	//!Notifying it's successor about it
	cn.successor.RpcNotify(cn.SelfNode) //!This will be a RPC Call
}

func (ch *ChordNode) notify(n_prime Node) {
	if ch.predecessorStatus == false || IsIdBetweenRangeRightEndExclusive(n_prime.node_id, ch.predecessor.node_id, ch.SelfNode.node_id) == true {
		ch.predecessor = n_prime
	}
}

func (ch *ChordNode) checkPredecessor() {
	isAlive := ch.SelfNode.RpcIsAlive(ch.predecessor)
	if isAlive == false { //!Mean node has failed
		ch.predecessorStatus = false
	}
}

func (cn *ChordNode) fixFingers() {

	cn.fingerTable.next = cn.fingerTable.next + 1

	if cn.fingerTable.next >= m {
		cn.fingerTable.next = 0
	}

	//!Generate hash Id for the number
	var int_hash_id HashId //!Write a message to fetch this Id
	int_hash_id = GenerateHashIdForFingerIndex(cn.SelfNode.node_id, cn.fingerTable.next)

	newSuccessorReturned := cn.findSuccessor(int_hash_id)
	cn.fingerTable.table[cn.fingerTable.next] = newSuccessorReturned
}

//!Write a function which checks what timer has went off and then do as instructed

// !This will be infinitely running
func (ch *ChordNode) perodicallyCheck() {
	//!Create three timers for stablization, fix fingers and check_predecessor
	stable_ticker := time.NewTicker(5 * time.Millisecond)
	check_p_ticker := time.NewTicker(4 * time.Second)
	fix_f_timer := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-stable_ticker.C:
			ch.stabilize()

		case <-check_p_ticker.C:
			ch.checkPredecessor()

		case <-fix_f_timer.C:
			ch.fixFingers()
		}
	}
}
