package c6p1

import (
	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func MaxValue(root tree.Node) int {
	maxVal := -1
	maxVal = max(maxVal, root.Data)
	if root.Left != nil {
		maxVal = max(maxVal, MaxValue(*root.Left))
	}
	if root.Right != nil {
		maxVal = max(maxVal, MaxValue(*root.Right))
	}
	return maxVal
}
