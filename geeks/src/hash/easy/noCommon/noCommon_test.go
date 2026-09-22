package noCommon_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/noCommon"

	"fmt"
	"testing"
)

type tData struct {
	input1 []int
	input2 []int
	want   int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input1: []int{2, 3, 4, 5, 8},
			input2: []int{1, 2, 3, 4},
			want:   3,
		},
		tData{
			input1: []int{1, 2, 3, 4},
			input2: []int{5, 6, 7},
			want:   0,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", noCommon.Print(v.input1, v.input2))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v -> %v\n",
				output, v.input1, v.input2, v.want,
			)
			return
		}
	}
}
