package c6p47_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/c6p47"
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func TestLCA(t *testing.T) {
	root := &tree.Node{}
	ins := []int{
		4,
		2, 8,
		0, 0, 5, 9,
		0, 0, 0, 0,
		0, 7, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 6, 0, 0, 0, 0, 0,
	}
	root.BuildBtree(ins)

	ck := [][]int{
		[]int{4, 4, 2},
		[]int{4, 2, 8},
		[]int{8, 9, 6},
		[]int{5, 5, 6},
	}
	var lca int
	for _, v := range ck {
		lca = c6p47.LCA(root, v[1], v[2])
		if v[0] != lca {
			t.Errorf("Wrong %d: %d, %d", lca, v[1], v[2])
		}
	}

}
