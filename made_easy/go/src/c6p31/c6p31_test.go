package c6p31_test

import (
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/c6p31"
	"github.com/Bozar/data-structures-and-algorithms/made_easy/go/src/tree"
)

func TestVerticalSum(t *testing.T) {
	root1 := tree.Node{}
	root1.BuildBtree([]int{
		1, 2, 3,
		4, 5, 6, 7,
	})
	sum1 := []int{4, 2, 12, 3, 7}

	root2 := tree.Node{}
	root2.BuildBtree([]int{
		1, 2, 3,
		4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15,
	})
	sum2 := []int{8, 4, 33, 12, 41, 7, 15}

	check := c6p31.VerticalSum(root1)
	if len(check) != len(sum1) {
		t.Errorf("Wrong len(check): %d", len(check))
		return
	}
	for k, v := range sum1 {
		if v != check[k] {
			t.Errorf("[%d] Wrong: %d, Right: %d", k, check[k], v)
		}
	}

	check = c6p31.VerticalSum(root2)
	if len(check) != len(sum2) {
		t.Errorf("Wrong len(check): %d", len(check))
		return
	}
	for k, v := range sum2 {
		if v != check[k] {
			t.Errorf("[%d] Wrong: %d, Right: %d", k, check[k], v)
		}
	}
}
