package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	q := New()
	q.Push("123")
	q.Push("234")
	q.Push("345")
	// check size
	if q.Len() != 3 {
		t.Fatal("size is wrong")
	}
	t1 := q.Pop()
	t2 := q.Pop()
	t3 := q.Peek()
	t4 := q.Pop()
	t.Log(t1, t2, t3, t4)
	if t1 != "345" {
		t.Fatal("wrong data")
	}

	if t2 != "234" {
		t.Fatal("wrong data")
	}

	if t3 != "123" {
		t.Fatal("wrong data")
	}

	if t4 != "123" {
		t.Fatal("wrong data")
	}
	if q.Len() != 0 {
		t.Fatal("wrong size")
	}
}
