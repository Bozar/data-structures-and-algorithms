package rotate_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/rotate"

	"fmt"
	"testing"
)

type tData struct {
	input    []int
	distance int
	want     []int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input:    []int{1, 2, 3, 4, 5, 6},
			want:     []int{5, 6, 1, 2, 3, 4},
			distance: 2,
		},
		tData{
			input:    []int{1, 2, 3},
			want:     []int{3, 1, 2},
			distance: 4,
		},
	}

	for _, v := range data {
		sOutput := fmt.Sprintf("%v", rotate.Print(v.input, v.distance))
		sWant := fmt.Sprintf("%v", v.want)
		if sWant != sOutput {
			t.Errorf("\nwrong: %s\nwant: %s\n", sOutput, sWant)
			return
		}
	}
}
