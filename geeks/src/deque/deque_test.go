package deque_test

import (
	"fmt"
	"github.com/Bozar/data-structures-and-algorithms/geeks/src/deque"
	"testing"
)

func TestDeque(t *testing.T) {
	fmt.Println("testNew()")
	testNew(t)
	testFront(t)
	testBack(t)
}

func testNew(t *testing.T) {
	dq := deque.New[int]()
	if !dq.IsEmpty() {
		t.Error("deque is NOT empty")
	} else if dq.Size() != 0 {
		t.Error("deque size is NOT 0")
	}
}

func testFront(t *testing.T) {
	dq := deque.New[int]()
	for i := 0; i < 10; i++ {
		dq.PushFront(i)
		if s := dq.Size(); s != i+1 {
			t.Errorf("wrong size: %d, want: %d", s, i+1)
		} else if e, _ := dq.PeekFront(); e != i {
			t.Errorf("wrong PeekFront: %d, want: %d", e, i+1)
		}
	}
	for i := 9; i > -1; i-- {
		if e, _ := dq.PopFront(); e != i {
			t.Errorf("wrong PopFront: %d, want: %d", e, i)
		} else if s := dq.Size(); s != i {
			t.Errorf("wrong size: %d, want: %d", s, i)
		}
	}
	if e, ok := dq.PopFront(); ok {
		t.Errorf("PopFront should fail: %d", e)
	}
}

func testBack(t *testing.T) {
	dq := deque.New[int]()
	for i := 0; i < 10; i++ {
		dq.PushBack(i)
		if s := dq.Size(); s != i+1 {
			t.Errorf("wrong size: %d, want: %d", s, i+1)
		} else if e, _ := dq.PeekBack(); e != i {
			t.Errorf("wrong PeekBack: %d, want: %d", e, i+1)
		}
	}
	for i := 9; i > -1; i-- {
		if e, _ := dq.PopBack(); e != i {
			t.Errorf("wrong PopBack: %d, want: %d", e, i)
		} else if s := dq.Size(); s != i {
			t.Errorf("wrong size: %d, want: %d", s, i)
		}
	}
	if e, ok := dq.PopBack(); ok {
		t.Errorf("PopBack should fail: %d", e)
	}
}
