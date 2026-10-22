package stack_test

import (
	"fmt"
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/geeks/src/stack"
)

func TestStack(t *testing.T) {
	fmt.Println("testNew()")
	testNew(t)

	fmt.Println()
	fmt.Println("testPush()")
	testPush(t)

	fmt.Println()
	fmt.Println("testPop()")
	testPop(t)
}

func testNew(t *testing.T) {
	stack := stack.New[int]()
	if !stack.IsEmpty() {
		t.Error("stack is NOT empty")
	} else if stack.Size() != 0 {
		t.Error("stack size is NOT 0")
	}
}

func testPush(t *testing.T) {
	stack := stack.New[int]()
	data, ok := stack.Peek()
	if ok {
		t.Errorf("Peek() should fail: %d", data)
	}

	for i := 1; i < 6; i++ {
		stack.Push(i)
		if stack.Size() != i {
			t.Errorf("wrong size: %d, right: %d", stack.Size(), i)
		}
		data, ok = stack.Peek()
		if (data != i) || (!ok) {
			t.Errorf("wrong peek data: %d, right: %d", data, i)
		}
	}
}

func testPop(t *testing.T) {
	stack := stack.New[int]()
	data, ok := stack.Pop()
	if ok {
		t.Errorf("Pop() should fail: %d", data)
	}

	const maxI = 5
	for i := maxI; i > 0; i-- {
		stack.Push(i)
	}
	for i := 1; i < maxI+1; i++ {
		data, ok = stack.Pop()
		if (data != i) || (!ok) {
			t.Errorf("wrong pop: %d, right: %d", data, i)
		} else if stack.Size() != maxI-i {
			t.Errorf(
				"wrong size: %d, right: %d",
				stack.Size(), maxI-i,
			)
		}
		if i == maxI {
			continue
		}
		data, ok = stack.Peek()
		if (data != i+1) || (!ok) {
			t.Errorf("wrong peek: %d, right: %d", data, i+1)
		}
	}
}
