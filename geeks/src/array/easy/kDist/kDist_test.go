package kDist_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/easy/kDist"

	"fmt"
	"testing"
)

type tData struct {
	input []int
	k     int
	want  bool
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{1, 2, 3, 4, 1, 2, 3, 4},
			k:     3,
			want:  false,
		},
		tData{
			input: []int{1, 2, 3, 1, 4, 5},
			k:     3,
			want:  true,
		},
		tData{
			input: []int{1, 2, 3, 4, 5},
			k:     3,
			want:  false,
		},
	}

	for _, v := range data {
		output := kDist.Print(v.input, v.k)
		sWant := fmt.Sprintf("%v: %v, %v", v.want, v.input, v.k)
		if output != v.want {
			t.Errorf("\nwrong: %v\nwant: %s\n", output, sWant)
			return
		}
	}
}
