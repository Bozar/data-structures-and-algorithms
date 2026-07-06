package mvZero_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/mvZero"

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
			input: []int{1, 2, 0, 4, 3, 0, 5, 0},
			want:  []int{1, 2, 4, 3, 5, 0, 0, 0},
		},
		tData{
			input: []int{10, 20, 30},
			want:  []int{10, 20, 30},
		},
		tData{
			input: []int{0, 0},
			want:  []int{0, 0},
		},
	}

	for _, v := range data {
		sOutput := fmt.Sprintf("%v", mvZero.Print(v.input))
		sWant := fmt.Sprintf("%v", v.want)
		if sWant != sOutput {
			t.Errorf("\nwrong: %s\nwant: %s\n", sOutput, sWant)
			return
		}
	}
}
