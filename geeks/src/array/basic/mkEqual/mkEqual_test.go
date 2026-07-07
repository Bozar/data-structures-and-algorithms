package mkEqual_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/mkEqual"

	"fmt"
	"testing"
)

type tData struct {
	input []int
	k     int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{4, 7, 19, 16},
			k:     3,
			want:  10,
		},
		tData{
			input: []int{4, 4, 4, 4},
			k:     3,
			want:  0,
		},
		tData{
			input: []int{4, 2, 6, 8},
			k:     3,
			want:  -1,
		},
	}

	for _, v := range data {
		output := mkEqual.Print(v.input, v.k)
		sWant := fmt.Sprintf("%v: %v, %v", v.want, v.input, v.k)
		if output != v.want {
			t.Errorf("\nwrong: %d\nwant: %s\n", output, sWant)
			return
		}
	}
}
