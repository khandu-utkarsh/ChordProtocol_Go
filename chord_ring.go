package chord

func (n Node) create() ChordNode {

	var cn ChordNode

	cn.self_node = n
	cn.successor = n
	cn.predecessor = nil
	cn.fingerTable = nil

	return cn
}

func (n Node) join(n_prime Node) ChordNode {

	var cn ChordNode

	cn.self_node = n
	cn.predecessor = nil
	cn.fingerTable = nil

	cn.successor = n_prime.RPC_find_successor(n.node_id)
	return cn
}
