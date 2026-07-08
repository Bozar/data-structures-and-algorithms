package single_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/easy/single"

	//	"fmt"
	"testing"
)

type tData struct {
	input []int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{2, 3, 5, 4, 5, 3, 4},
			want:  2,
		},
		tData{
			input: []int{2, 2, 5, 5, 20, 30, 30},
			want:  20,
		},
	}

	for _, v := range data {
		output := single.Print(v.input)
		if output != v.want {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v\n",
				output, v.want, v.input,
			)
			return
		}
	}
}
