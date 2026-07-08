package evGreater_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/easy/evGreater"

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
			input: []int{1, 2, 2, 1},
			want:  []int{1, 2, 1, 2},
		},
		tData{
			input: []int{1, 3, 2},
			want:  []int{1, 3, 2},
		},
		tData{
			input: []int{3, 1, 5},
			want:  []int{1, 5, 3},
		},
		tData{
			input: []int{4, 3, 1, 5},
			want:  []int{3, 4, 1, 5},
		},
	}

	for _, v := range data {
		sOutput := fmt.Sprintf("%v", evGreater.Print(v.input))
		sWant := fmt.Sprintf("%v", v.want)
		if sOutput != sWant {
			t.Errorf("\nwrong: %v\nwant: %s\n", sOutput, sWant)
			return
		}
	}
}
