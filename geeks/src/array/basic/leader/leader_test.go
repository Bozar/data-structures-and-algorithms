package leader_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/leader"

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
			input: []int{16, 17, 4, 3, 5, 2},
			want:  []int{17, 5, 2},
		},
		tData{
			input: []int{1, 2, 3, 4, 5, 2},
			want:  []int{5, 2},
		},
	}

	for _, v := range data {
		sWant := fmt.Sprintf("%v", v.want)
		sOutput := fmt.Sprintf("%v", leader.Print(v.input))
		if sWant != sOutput {
			t.Errorf("\nwrong: %s\nwant: %s\n", sOutput, sWant)
			return
		}
	}
}
