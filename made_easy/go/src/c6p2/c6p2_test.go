package c6p2_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/c6p2"
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func TestMaxValue(t *testing.T) {
	var root *tree.Node
	var nodes []*tree.Node
	var node *tree.Node
	var maxVal int
	for i := 0; i < 10; i++ {
		root = &tree.Node{}
		nodes = root.RandBtree(int64(42+i), 5, 99)
		node = nodes[int(float64(len(nodes)*(i+1))*0.1)-1]
		node.Data = 100
		maxVal = c6p2.MaxValue(*root)
		if maxVal != 100 {
			t.Errorf("Wrong max: %d", maxVal)
		}
	}
}
