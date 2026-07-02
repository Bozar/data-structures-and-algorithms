package c6p21_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/c6p21"
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func TestHasPath(t *testing.T) {
	root := &tree.Node{}
	root.RandBtree(42, 4, 10)
	// root.Print()

	// 4, 5, 6, ..., 28
	tf := []bool{
		false, true, false, false, true,
		true, false, true, true, false,
		false, true, true, false, false,
		true, true, true, true, false,
		true, false, true, false, true,
	}
	for i := 4; i < 29; i++ {
		if c6p21.HasPath(root, i, 0) != tf[i-4] {
			t.Errorf("%d: %v\n", i, tf[i-4])
		}
	}
}
