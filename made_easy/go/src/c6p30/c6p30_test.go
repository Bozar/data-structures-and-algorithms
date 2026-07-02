package c6p30_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/c6p30"
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func TestZigzagTraverse(t *testing.T) {
	nodes := []tree.Node{}
	for i := 1; i < 8; i++ {
		nodes = append(nodes, tree.Node{})
		nodes[len(nodes)-1].Data = i
	}
	root := nodes[0]
	root.Left = &nodes[1]
	root.Right = &nodes[2]
	(*root.Left).Left = &nodes[3]
	(*root.Left).Right = &nodes[4]
	(*root.Right).Left = &nodes[5]
	(*root.Right).Right = &nodes[6]
	//	root.Print()

	check := []int{
		1, 3, 2, 4, 5, 6, 7,
	}
	output := c6p30.ZigzagTraverse(root)
	if len(output) != len(check) {
		t.Errorf("Len: %d-%d", len(check), len(output))
		return
	}
	for k, v := range check {
		if v != output[k] {
			t.Errorf("Idx: %d, Val: %d", k, output[k])
			return
		}
	}
}
