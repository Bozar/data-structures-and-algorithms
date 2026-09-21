package missMinMax_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/missMinMax"

	"fmt"
	"testing"
)

type tData struct {
	input []int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{4, 5, 3, 8, 6},
			want:  1,
		},
		tData{
			input: []int{2, 1, 3},
			want:  0,
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", missMinMax.Print(v.input))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v -> %v\n",
				output, v.input, v.want,
			)
			return
		}
	}
}
