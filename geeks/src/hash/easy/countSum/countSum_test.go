package countSum_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/countSum"

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
			input1: []int{1, 5, 7, -1, 5},
			input2: 6,
			want:   3,
		},
		tData{
			input1: []int{1, 1, 1, 1},
			input2: 2,
			want:   6,
		},
		tData{
			input1: []int{10, 12, 10, 15, -1},
			input2: 125,
			want:   0,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", countSum.Print(v.input1, v.input2))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v -> %v\n",
				output, v.input1, v.input2, v.want,
			)
			return
		}
	}
}
