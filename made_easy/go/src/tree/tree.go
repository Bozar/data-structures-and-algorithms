package tree

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	printTree([]*Node{n}, 0)
}

func (n *Node) Slice() []int {
	const null int = -1
	queue := []*Node{n}
	output := []int{}
	var front *Node

	for len(queue) > 0 {
		front = queue[0]
		if front == nil {
			output = append(output, null)
		} else {
			output = append(output, front.Data)
			queue = append(queue, front.Left)
			queue = append(queue, front.Right)
		}
		queue = queue[1:]
	}
	return output
}

func (n *Node) RandBtree(seed int64, height int, maxData int) []*Node {
	maxData = max(1, maxData)

	rd := rand.New(rand.NewSource(seed))
	n.Data = rd.Intn(maxData)
	treeNodes := make([]*Node, 0, 8)
	treeNodes = append(treeNodes, n)
	nextLine := []*Node{n}
	var nextNodes []*Node

	for i := 0; i < height; i++ {
		nextLine, nextNodes = subBtree(nextLine, rd, maxData)
		treeNodes = append(treeNodes, nextNodes...)
	}
	return treeNodes
}

func (n *Node) BuildBtree(treeData []int) {
	nodes := []*Node{n}
	for _, v := range treeData {
		nodes[0].Data = v
		nodes[0].Left = &Node{}
		nodes[0].Right = &Node{}
		nodes = append(nodes, nodes[0].Left)
		nodes = append(nodes, nodes[0].Right)
		nodes = nodes[1:]
	}
	pruneBtree(n)
}

func (n *Node) InsertBST(data int) {
	if n == nil {
		return
	}

	if data > n.Data {
		if n.Right == nil {
			n.Right = &Node{}
			(*n.Right).Data = data
			return
		}
		(n.Right).InsertBST(data)
		return
	}

	if n.Left == nil {
		n.Left = &Node{}
		(*n.Left).Data = data
		return
	}
	(n.Left).InsertBST(data)
}

func (n *Node) DeleteBST(data int) *Node {
	if n == nil {
		return n
	} else if data > n.Data {
		n.Right = (n.Right).DeleteBST(data)
		return n
	} else if data < n.Data {
		n.Left = (n.Left).DeleteBST(data)
		return n
	}

	if n.Left == nil {
		n = n.Right
		return n
	} else if n.Right == nil {
		n = n.Left
		return n
	}
	lMax := (n.Left).MaxBST()
	n.Data = lMax.Data
	n.Left = (n.Left).DeleteBST(lMax.Data)
	return n
}

func (n *Node) BSTnode(data int) *Node {
	if n == nil {
		return nil
	}
	if n.Data == data {
		return n
	} else if n.Data < data {
		return (n.Right).BSTnode(data)
	} else {
		return (n.Left).BSTnode(data)
	}
}

func (n *Node) MaxBST() *Node {
	if n == nil {
		return nil
	}

	if n.Right == nil {
		return n
	}
	return (n.Right).MaxBST()
}

func (n *Node) MinBST() *Node {
	if n == nil {
		return nil
	}

	if n.Left == nil {
		return n
	}
	return (n.Left).MinBST()
}

func printTree(nodes []*Node, index int) {
	isLastLine := true
	nextLine := make([]*Node, 0)
	fmt.Printf("%d:\t", index)
	for i, v := range nodes {
		if v == nil {
			fmt.Printf("-- ")
			nextLine = append(nextLine, nil)
			nextLine = append(nextLine, nil)
		} else {
			fmt.Printf("%02d ", v.Data)
			nextLine = append(nextLine, v.Left)
			nextLine = append(nextLine, v.Right)
			isLastLine = false
		}
		if i%2 != 0 {
			fmt.Printf("| ")
		}
	}
	fmt.Printf("\n")
	if isLastLine {
		return
	}
	printTree(nextLine, index+1)
}

func subBtree(nodes []*Node, rd *rand.Rand, maxData int) ([]*Node, []*Node) {
	nextLine := make([]*Node, 0)
	nextNodes := make([]*Node, 0)

	for i, v := range nodes {
		if v == nil {
			nextLine = append(nextLine, nil)
			nextLine = append(nextLine, nil)
			continue
		}

		if i == 0 {
			v.Left = &Node{}
			(v.Left).Data = rd.Intn(maxData)
			nextNodes = append(nextNodes, v.Left)
			if rd.Intn(10)%3 != 0 {
				v.Right = &Node{}
				(v.Right).Data = rd.Intn(maxData)
				nextNodes = append(nextNodes, v.Right)
			}
			nextLine = append(nextLine, v.Left)
			nextLine = append(nextLine, v.Right)
			continue
		}

		if rd.Intn(10)%3 != 0 {
			v.Left = &Node{}
			(v.Left).Data = rd.Intn(maxData)
			nextNodes = append(nextNodes, v.Left)
		}
		if rd.Intn(10)%3 != 0 {
			v.Right = &Node{}
			(v.Right).Data = rd.Intn(maxData)
			nextNodes = append(nextNodes, v.Right)
		}
		nextLine = append(nextLine, v.Left)
		nextLine = append(nextLine, v.Right)
	}
	return nextLine, nextNodes
}

func pruneBtree(root *Node) *Node {
	if (root == nil) || (root.Data == 0) {
		return nil
	}
	root.Left = pruneBtree(root.Left)
	root.Right = pruneBtree(root.Right)
	return root
}
