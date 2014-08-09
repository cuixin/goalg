package hlist

import (
	"fmt"
	"testing"
)

func TestHlist(t *testing.T) {
	l := New()
	l.PushFront(1)
	l.PushFront(2)
	first := l.Front()
	if first == nil {
		t.Fatal("first can not be nil")
	}
	if i, ok := first.Value.(int); !ok {
		t.Fatal("first.Value assection failed")
	} else {
		if i != 2 {
			t.Fatalf("i value error: %d", i)
		}
	}
	if next := first.Next(); next == nil {
		t.Fatal("next can not be nil")
	} else {
		if i, ok := next.Value.(int); !ok {
			t.Fatal("next.Value assection failed")
		} else {
			if i != 1 {
				t.Fatalf("i value error: %d", i)
			}
		}
	}
	if l.Len() != 2 {
		t.Fatal("length error")
	}
	l.PushFront(3)
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(6)
	first = l.Front()
	if l.Len() != 6 {
		t.Fatal("length error")
	}
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); !ok {
			t.Fatal("e.Value assection failed")
		} else {
			fmt.Println(i)
		}
	}
	fmt.Println("------")
	if i, ok := l.Remove(first).(int); !ok {
		t.Fatal("first.Value assection failed")
	} else {
		if i != 6 {
			t.Fatalf("i value error: %d", i)
		}
	}
	if l.Len() != 5 {
		t.Fatalf("length error")
	}
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); !ok {
			t.Fatal("e.Value assection failed")
		} else {
			fmt.Println(i)
		}
	}
	second := l.Front()    // 5
	thrid := second.Next() // 4
	fourth := thrid.Next() // 3
	fifth := fourth.Next() // 2
	sixth := fifth.Next()  // 1
	l.Remove(second)
	l.Remove(thrid)
	l.Remove(fourth)
	l.Remove(fifth)
	l.Remove(sixth)
	if l.Len() != 0 {
		t.Fatalf("length error")
	}
	first = l.Front()
	if first != nil {
		t.Fatal("first != nil")
	}
	for e := l.Front(); e != nil; e = e.Next() {
		if i, ok := e.Value.(int); !ok {
			t.Fatal("e.Value assection failed")
		} else {
			fmt.Println(i)
		}
	}

	e := l.PushFront(7)
	if i, ok := e.Value.(int); !ok {
		t.Fatal("e.Value assection failed")
	} else {
		if i != 7 {
			t.Fatalf("i value error: %d", i)
		}
	}
}
