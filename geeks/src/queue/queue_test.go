package queue_test

import (
	"fmt"
	"testing"

	"github.com/Bozar/data-structures-and-algorithms/geeks/src/queue"
)

func TestQueue(t *testing.T) {
	fmt.Println("testNew()")
	testNew(t)

	fmt.Println("testEnqueue()")
	testEnqueue(t)

	fmt.Println("testDequeue()")
	testDequeue(t)
}

func testNew(t *testing.T) {
	queue := queue.New[int]()
	if !queue.IsEmpty() {
		t.Error("queue is NOT empty")
	} else if queue.Size() != 0 {
		t.Error("queue size is NOT 0")
	}
}

func testEnqueue(t *testing.T) {
	queue := queue.New[int]()

	data, ok := queue.Peek()
	if ok {
		t.Errorf("Peek() should fail: %d", data)
	}

	const minI int = 1
	const maxI int = 6
	for i := minI; i < maxI; i++ {
		queue.Enqueue(i)
		if queue.Size() != i {
			t.Errorf("wrong size: %d, right: %d", queue.Size(), i)
		}
		data, ok = queue.Peek()
		if (data != minI) || (!ok) {
			t.Errorf("wrong peek data: %d, right: %d", data, minI)
		}
	}
}

func testDequeue(t *testing.T) {
	queue := queue.New[int]()

	data, ok := queue.Dequeue()
	if ok {
		t.Errorf("Dequeue() should fail: %d", data)
	}

	const minI int = 1
	const maxI int = 5
	for i := minI; i < maxI+1; i++ {
		queue.Enqueue(i)
	}
	for i := minI; i < maxI+1; i++ {
		data, ok = queue.Dequeue()
		if (data != i) || (!ok) {
			t.Errorf("wrong dequeue: %d, right: %d", data, i)
		} else if queue.Size() != maxI-i {
			t.Errorf(
				"wrong size: %d, right: %d",
				queue.Size(), maxI-i,
			)
		}
		if i == maxI {
			continue
		}
		data, ok = queue.Peek()
		if (data != i+1) || (!ok) {
			t.Errorf("wrong peek: %d, right: %d", data, i+1)
		}
	}
}
