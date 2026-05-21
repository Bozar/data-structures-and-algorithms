package c6p30

import (
	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func ZigzagTraverse(root tree.Node) []int {
	output := []int{}
	output = append(output, zigzag([]tree.Node{root}, true)...)
	return output
}

func zigzag(nodes []tree.Node, l2r bool) []int {
	output := []int{}
	if len(nodes) == 0 {
		return output
	}

	nextLine := []tree.Node{}
	for i := len(nodes) - 1; i > -1; i-- {
		output = append(output, nodes[i].Data)
		if (nodes[i].Left == nil) || (nodes[i].Right == nil) {
			continue
		}
		if l2r {
			nextLine = append(nextLine, *nodes[i].Left)
			nextLine = append(nextLine, *nodes[i].Right)
		} else {
			nextLine = append(nextLine, *nodes[i].Right)
			nextLine = append(nextLine, *nodes[i].Left)
		}
	}
	output = append(output, zigzag(nextLine, !l2r)...)
	return output
}
