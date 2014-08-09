package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	q.Enqueue("123")
	q.Enqueue("234")
	q.Enqueue("345")
	// check size
	if q.Len() != 3 {
		t.Fatal("size is wrong")
	}
	t1 := q.Dequeue()
	t2 := q.Dequeue()
	t3 := q.Peek()
	t4 := q.Dequeue()
	t.Log(t1, t2, t3, t4)
	if t1 != "123" {
		t.Fatal("wrong data")
	}

	if t2 != "234" {
		t.Fatal("wrong data")
	}

	if t3 != "345" {
		t.Fatal("wrong data")
	}

	if t4 != "345" {
		t.Fatal("wrong data")
	}
	if q.Len() != 0 {
		t.Fatal("wrong size")
	}
}
