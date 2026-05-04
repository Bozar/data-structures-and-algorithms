package c6p2

import (
	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func MaxValue(root tree.Node) int {
	maxVal := -1
	nodes := []tree.Node{root}
	var node tree.Node
	for len(nodes) > 0 {
		node = nodes[0]
		maxVal = max(maxVal, node.Data)
		nodes = nodes[1:]
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return maxVal
}
