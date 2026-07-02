package c6p47

import (
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func LCA(root *tree.Node, this int, that int) int {
	if ((this <= root.Data) && (that >= root.Data)) ||
		((this >= root.Data) && (that <= root.Data)) {
		return root.Data
	} else if this <= root.Data {
		return LCA(root.Left, this, that)
	} else {
		return LCA(root.Right, this, that)
	}
}
