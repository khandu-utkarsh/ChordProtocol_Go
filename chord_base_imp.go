package chord

// Constants
const m = 6

type HashId struct {
	id []byte
}

type Key struct {
	hash_val HashId
}

// ! Each process on the system will be identified by type Node -  abstraction over Id and port number and hashId for that process
type Node struct {
	node_id          HashId
	node_ip_address  string
	node_port_number string
}

// Finger type denoting identifying information about a ChordNode
type FingerTable struct {
	table [m]Node //m entities
	next  int

	//!In paper all the indexing for the entries in the table have been starting from 1, and in go indexing starts from 0, so whatever operations we are doing, simply
	// add 1 else it is a tedious process to change paper's pseudo code

	//Hence following should be true
	//table[0] Succeeds curr Node id by atleast 2^(0 + 1 -1) or curr_node_id + 2^(0 + 1 -1)
	//table[1] Succeeds curr Node id by atleast 2^(1 + 1 -1) or curr_node_id + 2^(1 + 1 -1)
	//table[2] Succeeds curr Node id by atleast 2^(2 + 1 -1) or curr_node_id + 2^(2 + 1 -1)
	//table[3] Succeeds curr Node id by atleast 2^(3 + 1 -1) or curr_node_id + 2^(3 + 1 -1)
	//table[4] Succeeds curr Node id by atleast 2^(4 + 1 -1) or curr_node_id + 2^(4 + 1 -1)
	//.
	//.
	//.
	//table[m -1 ] Succeeds curr Node id by atleast 2^(m-1 + 1 -1) or curr_node_id + 2^(5 + 1 -1)
}

//!Every machine (processor) will be of type ChordNode
//!It can be imagined as on the system each node will identity as ChordNode
//!So when a processor joins the node ring or even creates a new ring, a ChordNode object for that will be returned

// !Cord Node will be basically current node (current process)
type ChordNode struct {
	fingerTable FingerTable
	self_node   Node

	successor   Node //Assign this to the pointer to the first node of the table
	predecessor Node
}

func (cn ChordNode) closest_preceding_node(key Key) Node {
	for i := m - 1; i >= 0; i-- {
		if Is_X_BetweenRange_REExclusive(cn.fingerTable.table[i].node_id, cn.self_node.node_id, key.hash_val) {
			return cn.fingerTable.table[i]
		}
	}
	return cn.node_id
}

func (ch ChordNode) find_predecessor(key Key) Node {
	n_prime := ch.self_node
	n_prime_succ := n_prime.RPC_get_successor_node()

	for Is_X_BetweenRange_REInclusive(key.hash_val, n_prime.node_id, n_prime_succ.node_id) == false {
		n_prime = n_prime.RPC_closest_preceeding_node()
	}
	return n_prime
}

func (cn ChordNode) find_successor(key Key) Node {
	if Is_X_BetweenRange_REInclusive(key.hash_val, n_prime.node_id, cn.successor.node_id) {
		return cn.successor
	} else {
		n_prime = cn.closest_preceding_node(key)
		return n_prime.RPC_get_successor_node()
	}
}

//!Following three functions are called periodically
//	1. stabilize
//	2. fix_fingers
//	3. check_predecessor

func (ch ChordNode) stabilize() {
	succ_pre_node := ch.successor.RPC_GetPredecessor()
	if Is_X_BetweenRange_REExclusive(succ_pre_node.node_id, cn.self_node.node_id, cn.successor.node_id) {
		ch.successor = succ_pre_node //!Changing it successor
	}
	//!Notifying it's successor about it
	ch.successor.RPC_notify(ch.self_node) //!This will be a RPC Call
}

func (ch ChordNode) notify(n_prime Node) {
	if ch.predecessor == nil || Is_X_BetweenRange_REExclusive(n_prime.node_id, ch.predecessor.node_id, ch.self_node.node_id) == true {
		ch.predecessor = n_prime
	}
}

// Fix fingers also runs periodically
func (ch ChordNode) fix_fingers() {
	//Generate an random number between 0 to m-1 (We have m entries in the table)
	if ch.fingerTable.next > m-1 {
		ch.fingerTable.next = 0
	}
	key := Key
	key.hash_val = ch.self_node.node_id + pow(2, ch.fingerTable.next+1-1)
	ch.fingerTable.table[ch.fingerTable.next] = ch.find_successor(key)
}

func (ch ChordNode) check_predecessor() {
	if ch.predecessor.RPC_HasFailed() == true {
		ch.predecessor = nil
	}
}
