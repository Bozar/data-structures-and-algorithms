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
