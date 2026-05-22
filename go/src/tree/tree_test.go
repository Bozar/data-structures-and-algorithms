package tree_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func TestPrint(t *testing.T) {
	n := tree.Node{}
	n1 := tree.Node{}
	n2 := tree.Node{}
	n3 := tree.Node{}

	n.Data = 1
	n1.Data = 2
	n2.Data = 3
	n3.Data = 4

	n.Left = &n1
	n.Right = &n2
	n2.Left = &n3

	n.Print()
}

func TestRandBtree(t *testing.T) {
	n := tree.Node{}
	n.RandBtree(42, 4, 99)
	n.Print()
}

func TestBuildBtree(t *testing.T) {
	n := tree.Node{}
	n.BuildBtree([]int{1, 2, 0, 3, 4, 0, 0, 0, 5})
	n.Print()
}
