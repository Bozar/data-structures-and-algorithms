package stock_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/easy/stock"

	//	"fmt"
	"testing"
)

type tData struct {
	input []int
	want  int
}

func TestPrint(t *testing.T) {
	data := []tData{
		tData{
			input: []int{100, 180, 260, 310, 40, 535, 695},
			want:  865,
		},
		tData{
			input: []int{4, 2},
			want:  0,
		},
	}

	for _, v := range data {
		output := stock.Print(v.input)
		if output != v.want {
			t.Errorf(
				"\nwrong: %v\nwant: %v, %v\n",
				output, v.want, v.input,
			)
			return
		}
	}
}
