package rmvDup_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/rmvDup"

	"fmt"
	"testing"
)

type tData struct {
	input []int
	want  []int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{2, 2, 2, 2, 2},
			want:  []int{2},
		},
		tData{
			input: []int{1, 2, 2, 3, 4, 4, 4, 5, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		tData{
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, v := range data {
		sWant := fmt.Sprintf("%v", v.want)
		sOutput := fmt.Sprintf("%v", rmvDup.Print(v.input))
		if sWant != sOutput {
			t.Errorf("\nwrong: %s\nwant: %s\n", sOutput, sWant)
			return
		}
	}
}
