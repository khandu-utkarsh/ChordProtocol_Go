package chord

func (node Node) RPC_find_successor(id HashId) Node {

	//node should call find_successor funcation on it's machine and get this value
	var nodeReturned Node
	return nodeReturned
}

func (node Node) RPC_find_predecessor() Node {
	var nodeReturned Node
	return nodeReturned
}

func (send_node Node) RPC_notify(receive_node Node) bool {

	//!node should call notify on it with n_prime as argument
	successs := true
	return successs
}

func (send_node Node) RPC_IsAlive(receive_node Node) bool {
	//! send_node should send messages to receive node to check if it is alive or not,
	// if yes, return ture and if not return false
	return false
}