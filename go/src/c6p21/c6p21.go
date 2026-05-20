package c6p21

import (
	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func HasPath(root *tree.Node, sum int, parent int) bool {
	if root == nil {
		return false
	}

	rd := root.Data + parent
	if rd > sum {
		return false
	} else if rd == sum {
		return true
	}
	return HasPath(root.Left, sum, rd) ||
		HasPath(root.Right, sum, rd)
}
