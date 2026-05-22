package c6p31

import (
	"sort"

	"github.com/Bozar/data-structures-and-algorithms-made-easy/go/src/tree"
)

func VerticalSum(root tree.Node) []int {
	mSum := map[int]int{}
	getSum(root, mSum, 0)

	keys := []int{}
	for k := range mSum {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	sSum := []int{}
	for _, v := range keys {
		sSum = append(sSum, mSum[v])
	}
	return sSum
}

func getSum(root tree.Node, sum map[int]int, column int) {
	if root.Left != nil {
		getSum(*root.Left, sum, column-1)
	}
	sum[column] += root.Data
	if root.Right != nil {
		getSum(*root.Right, sum, column+1)
	}
}
