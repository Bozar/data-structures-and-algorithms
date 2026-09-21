package missRange_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/hash/easy/missRange"

	"fmt"
	"testing"
)

type tData struct {
	input1 []int
	input2 int
	input3 int
	want   []int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input1: []int{10, 12, 11, 15},
			input2: 10,
			input3: 15,
			want:   []int{13, 14},
		},
		tData{
			input1: []int{1, 4, 11, 51, 15},
			input2: 50,
			input3: 55,
			want:   []int{50, 52, 53, 54, 55},
		},
	}

	for _, v := range data {
		output := fmt.Sprintf("%v", missRange.Print(
			v.input1, v.input2, v.input3,
		))
		if output != fmt.Sprintf("%v", v.want) {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v, %v -> %v\n",
				output, v.input1, v.input2, v.input3, v.want,
			)
			return
		}
	}
}
