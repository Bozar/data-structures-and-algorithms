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

func TestSlice(t *testing.T) {
	//  1
	// 2 3
	r1 := &tree.Node{}
	r1.Data = 1
	r1.Left = &tree.Node{}
	(*r1.Left).Data = 2
	r1.Right = &tree.Node{}
	(*r1.Right).Data = 3
	s1 := r1.Slice()
	c1 := []int{
		1,
		2, 3,
		-1, -1, -1, -1,
	}
	if len(s1) != len(c1) {
		t.Errorf("Wrong len: %d", len(s1))
		return
	}
	for i, v := range c1 {
		if v != s1[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, s1[i], v)
		}
	}

	//   1
	//  2
	// 3
	r2 := &tree.Node{}
	r2.Data = 1
	r2.Left = &tree.Node{}
	(*r2.Left).Data = 2
	(*r2.Left).Left = &tree.Node{}
	(*(*r2.Left).Left).Data = 3
	s2 := r2.Slice()
	c2 := []int{
		1,
		2, -1,
		3, -1,
		-1, -1,
	}
	if len(s2) != len(c2) {
		t.Errorf("Wrong len: %d", len(s2))
		return
	}
	for i, v := range c2 {
		if v != s2[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, s2[i], v)
		}
	}

	//   1
	// 2
	//   3
	r3 := &tree.Node{}
	r3.Data = 1
	r3.Left = &tree.Node{}
	(*r3.Left).Data = 2
	(*r3.Left).Right = &tree.Node{}
	(*(*r3.Left).Right).Data = 3
	s3 := r3.Slice()
	c3 := []int{
		1,
		2, -1,
		-1, 3,
		-1, -1,
	}
	if len(s3) != len(c3) {
		t.Errorf("Wrong len: %d", len(s3))
		return
	}
	for i, v := range c3 {
		if v != s3[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, s3[i], v)
		}
	}
}

func TestInsertBST(t *testing.T) {
	// 1
	//   2
	//     3
	//       4
	//         5
	r1 := &tree.Node{}
	ins1 := []int{1, 2, 3, 4, 5}
	for i, v := range ins1 {
		if i == 0 {
			r1.Data = v
			continue
		}
		r1.InsertBST(v)
	}
	sl1 := r1.Slice()
	ck1 := []int{
		1,
		-1, 2,
		-1, 3,
		-1, 4,
		-1, 5,
		-1, -1,
	}
	if len(sl1) != len(ck1) {
		t.Errorf("Wrong len: %d", len(sl1))
		return
	}
	for i, v := range ck1 {
		if v != sl1[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl1[i], v)
		}
	}

	// 1
	//   3
	// 2   5
	//    4
	r2 := &tree.Node{}
	ins2 := []int{1, 3, 5, 2, 4}
	for i, v := range ins2 {
		if i == 0 {
			r2.Data = v
			continue
		}
		r2.InsertBST(v)
	}
	sl2 := r2.Slice()
	ck2 := []int{
		1,
		-1, 3,
		2, 5,
		-1, -1, 4, -1,
		-1, -1,
	}
	if len(sl2) != len(ck2) {
		t.Errorf("Wrong len: %d", len(sl2))
		return
	}
	for i, v := range ck2 {
		if v != sl2[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl2[i], v)
		}
	}

	//    3
	//  2   5
	// 1   4
	r3 := &tree.Node{}
	ins3 := []int{3, 2, 5, 1, 4}
	for i, v := range ins3 {
		if i == 0 {
			r3.Data = v
			continue
		}
		r3.InsertBST(v)
	}
	sl3 := r3.Slice()
	ck3 := []int{
		3,
		2, 5,
		1, -1, 4, -1,
		-1, -1, -1, -1,
	}
	if len(sl3) != len(ck3) {
		t.Errorf("Wrong len: %d", len(sl3))
		return
	}
	for i, v := range ck3 {
		if v != sl3[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl3[i], v)
		}
	}
}

func TestBSTnode(t *testing.T) {
	root := &tree.Node{}
	insert := []int{3, 2, 5, 1, 4}
	for i, v := range insert {
		if i == 0 {
			root.Data = v
			continue
		}
		root.InsertBST(v)
	}

	var node *tree.Node
	for i := 0; i < 7; i++ {
		node = root.BSTnode(i)
		if (i == 0) || (i == 6) {
			if node != nil {
				t.Errorf("Not nil: %d", i)
			}
		} else if (node == nil) || (node.Data != i) {
			t.Errorf("Is wrong node: %d", i)
		}
	}
}

func TestMaxBST(t *testing.T) {
	root := &tree.Node{}
	insert := []int{3, 2, 5, 1, 4}
	for i, v := range insert {
		if i == 0 {
			root.Data = v
			continue
		}
		root.InsertBST(v)
	}

	node := root.MaxBST()
	if node.Data != 5 {
		t.Errorf("Max: %d", node.Data)
	}
}

func TestMinBST(t *testing.T) {
	root := &tree.Node{}
	insert := []int{3, 2, 5, 1, 4}
	for i, v := range insert {
		if i == 0 {
			root.Data = v
			continue
		}
		root.InsertBST(v)
	}

	node := root.MinBST()
	if node.Data != 1 {
		t.Errorf("Min: %d", node.Data)
	}
}

func TestDeleteBST(t *testing.T) {
	// 1
	//   2
	//     3
	//       4
	//         5x
	r1 := &tree.Node{}
	ins1 := []int{1, 2, 3, 4, 5}
	for i, v := range ins1 {
		if i == 0 {
			r1.Data = v
			continue
		}
		r1.InsertBST(v)
	}
	r1.DeleteBST(5)
	sl1 := r1.Slice()
	ck1 := []int{
		1,
		-1, 2,
		-1, 3,
		-1, 4,
		-1, -1,
	}
	if len(sl1) != len(ck1) {
		t.Errorf("Wrong len: %d", len(sl1))
		return
	}
	for i, v := range ck1 {
		if v != sl1[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl1[i], v)
		}
	}

	// 1
	//   3x
	// 2   5
	//    4
	r2 := &tree.Node{}
	ins2 := []int{1, 3, 5, 2, 4}
	for i, v := range ins2 {
		if i == 0 {
			r2.Data = v
			continue
		}
		r2.InsertBST(v)
	}
	r2.DeleteBST(3)
	sl2 := r2.Slice()
	ck2 := []int{
		1,
		-1, 2,
		-1, 5,
		4, -1,
		-1, -1,
	}
	if len(sl2) != len(ck2) {
		t.Errorf("Wrong len: %d", len(sl2))
		return
	}
	for i, v := range ck2 {
		if v != sl2[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl2[i], v)
		}
	}

	//    3
	//  2   5
	// 1   4x
	r3 := &tree.Node{}
	ins3 := []int{3, 2, 5, 1, 4}
	for i, v := range ins3 {
		if i == 0 {
			r3.Data = v
			continue
		}
		r3.InsertBST(v)
	}
	r3.DeleteBST(4)
	sl3 := r3.Slice()
	ck3 := []int{
		3,
		2, 5,
		1, -1, -1, -1,
		-1, -1,
	}
	if len(sl3) != len(ck3) {
		t.Errorf("Wrong len: %d", len(sl3))
		return
	}
	for i, v := range ck3 {
		if v != sl3[i] {
			t.Errorf("[%d] Wrong: %d, Right: %d", i, sl3[i], v)
		}
	}
}
