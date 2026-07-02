package leader_test

import (
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/array/basic/leader"

	"testing"
)

func TestPrint(t *testing.T) {
	input := []int{
		16, 17, 4, 3, 5, 2,
	}
	want := []int{
		17, 5, 2,
	}
	test(input, want, t)

	input = []int{
		1, 2, 3, 4, 5, 2,
	}
	want = []int{
		5, 2,
	}
	test(input, want, t)
}

func test(input []int, want []int, t *testing.T) {
	output := leader.Print(input)
	lenWant := len(want)
	lenOutput := len(output)
	if lenWant != lenOutput {
		t.Errorf("wrong len: %d, want: %d\n", lenOutput, lenWant)
		return
	}
	for i, v := range want {
		if v != output[i] {
			t.Errorf("\nwrong: %v\nwant: %v\n", output, want)
			return
		}
	}
}
