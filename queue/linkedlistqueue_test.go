package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	queue := New()
	if queue.Head != nil && queue.Tail != nil && queue.length != 0 {
		t.Fail()
	}
	queue.Enqueue(1)
	if queue.Tail.Value != 1 {
		t.Fail()
	}
	queue.Enqueue(1)
	if queue.length != 2 {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	queue := New()
	if err := queue.Dequeue(); err == nil {
		t.Fail()
	}
	queue.Enqueue(5)
	queue.Enqueue(7)
	if err := queue.Dequeue(); err != nil {
		t.Fail()
	}
}