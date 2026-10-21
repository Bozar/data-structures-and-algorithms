package celebrity_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/stack/easy/celebrity"

	"fmt"
	"testing"
)

type tData struct {
	input [][]int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{0, 1, 1},
			},
			want: 1,
		},
		tData{
			input: [][]int{
				{1, 1},
				{1, 1},
			},
			want: -1,
		},
		tData{
			input: [][]int{
				{0},
			},
			want: 0,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", celebrity.Print(v.input))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v -> %v\n",
				output, v.input, v.want,
			)
			return
		}
	}
}
