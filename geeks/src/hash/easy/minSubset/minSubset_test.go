package minSubset_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/minSubset"

	"fmt"
	"testing"
)

type tData struct {
	input []int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{1, 2, 3, 4},
			want:  1,
		},
		tData{
			input: []int{1, 2, 3, 3},
			want:  2,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", minSubset.Print(v.input))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v -> %v\n",
				output, v.input, v.want,
			)
			return
		}
	}
}
