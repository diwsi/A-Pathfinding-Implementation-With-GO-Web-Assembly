package AStar

const EMPTY = 0
const START = 1
const TARGET = 2
const BLOCKED = 3

type Node struct {
	x         uint16
	y         uint16
	index     uint32
	val       uint8
	gCost     uint32
	hCost     uint32
	totalCost uint32
	open      bool
	checked   bool
	parent    *Node
}

func (n Node) Equals(node *Node) bool {
	return n.x == node.x && n.y == node.y
}
