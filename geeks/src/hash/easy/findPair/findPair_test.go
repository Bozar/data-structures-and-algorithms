package findPair_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/findPair"

	"fmt"
	"testing"
)

type tData struct {
	input1 []int
	input2 int
	want   bool
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input1: []int{0, -1, 2, -3, 1},
			input2: -2,
			want:   true,
		},
		tData{
			input1: []int{1, -2, 1, 0, 5},
			input2: 0,
			want:   false,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", findPair.Print(v.input1, v.input2))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v -> %v\n",
				output, v.input1, v.input2, v.want,
			)
			return
		}
	}
}
