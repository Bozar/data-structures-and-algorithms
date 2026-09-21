package countDiff_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/countDiff"

	"fmt"
	"testing"
)

type tData struct {
	input1 []int
	input2 int
	want   int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input1: []int{1, 4, 1, 4, 5},
			input2: 3,
			want:   4,
		},
		tData{
			input1: []int{8, 16, 12, 16, 4, 0},
			input2: 4,
			want:   5,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", countDiff.Print(v.input1, v.input2))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v -> %v\n",
				output, v.input1, v.input2, v.want,
			)
			return
		}
	}
}
