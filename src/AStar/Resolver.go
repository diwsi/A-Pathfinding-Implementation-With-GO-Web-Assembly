package AStar

func Resolve(problem [][]uint8) []uint32 {
	//map to node objects
	nodes := mapToNodes(problem)

	startNode := FindNode(nodes, START)
	endNode := FindNode(nodes, TARGET)
	targetNode := FindNode(nodes, START)

	var openNodes []*Node
	var solutionNodes []uint32

	for {
		nodes[targetNode.y][targetNode.x].checked = true
		adjentNodes := getAdjacentNodes(nodes, targetNode.x, targetNode.y)
		//Search new nodes
		for _, node := range adjentNodes {
			if node.open || node.checked {
				continue
			}
			node.parent = targetNode
			calculateCost(node, startNode, endNode)
			nodes[targetNode.y][targetNode.x].open = true
			openNodes = append(openNodes, node)
		}

		//New node to investigate
		targetNode = getLowestCostNode(openNodes)

		openNodes = remove(openNodes, targetNode)
		if !(targetNode != nil || len(openNodes) > 0) {
			//No valid nodes left to find solution
			break
		}

		//Solution Found
		if targetNode.Equals(endNode) {
			sNode := targetNode
			for {
				//Backtrace the path
				solutionNodes = append(solutionNodes, sNode.index)
				if sNode.parent.Equals(startNode) {
					break
				}
				sNode = sNode.parent
			}
			break
		}
	}
	return solutionNodes
}

func remove(nodes []*Node, node *Node) []*Node {
	//Remove node pointer
	for i, n := range nodes {
		if n == node {
			nodes = append(nodes[:i], nodes[i+1:]...)
			break
		}
	}
	return nodes
}

func getLowestCostNode(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	lowestCostNode := nodes[0]
	for _, node := range nodes {
		//lower hcost closer to target node
		if node.hCost <= lowestCostNode.hCost {
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func calculateCost(node *Node, start *Node, end *Node) {
	//Distance to start point
	node.gCost = Abs(int16(start.x)-int16(node.x)) + Abs(int16(start.y)-int16(node.y))
	//Distance to end point
	node.hCost = Abs(int16(end.x)-int16(node.x)) + Abs(int16(end.y)-int16(node.y))
	node.totalCost = node.gCost + node.hCost
}

func mapToNodes(set [][]uint8) [][]Node {
	var nodes [][]Node
	for i, row := range set {
		var nrow []Node
		for j, col := range row {
			//Create node for calculations
			nNode := Node{
				index: uint32(i*len(row) + j),
				val:   col,
				y:     uint16(i),
				x:     uint16(j),
			}
			nrow = append(nrow, nNode)
		}
		nodes = append(nodes, nrow)
	}
	return nodes
}

func getAdjacentNodes(source [][]Node, x uint16, y uint16) []*Node {
	var nodes []*Node
	//For way adjacent node maps left,right,down,up
	adjacenNodeMap := [][]int16{
		[]int16{-1, 0},
		[]int16{1, 0},
		[]int16{0, -1},
		[]int16{0, 1},
	}

	for _, mp := range adjacenNodeMap {
		refNode := getNode(source, uint16(int16(x)-mp[0]), uint16(int16(y)-mp[1]))
		if refNode != nil && refNode.val != BLOCKED { //skip blocked nodes
			nodes = append(nodes, refNode)
		}
	}

	return nodes
}

func getNode(source [][]Node, x uint16, y uint16) *Node {
	hLen := uint16(len(source))
	wLen := uint16(len(source[0]))

	if y < 0 || y >= hLen || x < 0 || x >= wLen {
		return nil
	}
	return &source[y][x] //node ref. by matrix coordinates
}

func FindNode(source [][]Node, nodeType uint8) *Node {
	for _, row := range source {
		for _, n := range row {
			if n.val == nodeType { //Find Node by type  value
				return &n
			}
		}
	}
	return nil
}

func Abs(n int16) uint32 {

	if n > 0 {
		return uint32(n)
	}
	return uint32(-n)

}
