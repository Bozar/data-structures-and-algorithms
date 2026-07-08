package sumSub_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/easy/sumSub"

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
			input: []int{1, 4, 5, 3, 2},
			want:  116,
		},
		tData{
			input: []int{1, 2, 3, 4},
			want:  50,
		},
	}

	for _, v := range data {
		output := sumSub.Print(v.input)
		if output != v.want {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v\n",
				output, v.want, v.input,
			)
			return
		}
	}
}
