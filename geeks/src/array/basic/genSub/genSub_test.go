package genSub_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/genSub"

	"fmt"
	"slices"
	"testing"
)

type tData struct {
	input []int
	want  [][]int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{1, 2, 3},
			want: [][]int{
				[]int{1}, []int{2}, []int{3},
				[]int{1, 2}, []int{2, 3},
				[]int{1, 2, 3},
			},
		},
		tData{
			input: []int{1, 2},
			want: [][]int{
				[]int{1}, []int{2},
				[]int{1, 2},
			},
		},
	}

	for _, v := range data {
		outIter := genSub.PrintIter(v.input)
		slices.SortStableFunc(outIter, sortOutput)
		outRec := genSub.PrintRec(v.input)
		slices.SortStableFunc(outRec, sortOutput)

		sWant := fmt.Sprintf("%v", v.want)
		sOutIter := fmt.Sprintf("%v", outIter)
		sOutRec := fmt.Sprintf("%v", outRec)
		if sWant != sOutIter {
			t.Errorf("\niterative: %s\nwant: %s\n", sOutIter, sWant)
			return
		} else if sWant != sOutRec {
			t.Errorf("\nrecursive: %s\nwant: %s\n", sOutRec, sWant)
			return
		}
	}
}

func sortOutput(a, b []int) int {
	if len(a) > len(b) {
		return 1
	} else if len(a) < len(b) {
		return -1
	}
	for i, v := range a {
		if v > b[i] {
			return 1
		} else if v < b[i] {
			return -1
		}
	}
	return 0
}
