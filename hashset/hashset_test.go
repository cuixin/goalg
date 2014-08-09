package hashset

import (
	"testing"
)

func TestHashSet(t *testing.T) {
	h := New()
	h.Add("123")
	h.Add("234")
	h.Add("345")
	// check size
	if h.Len() != 3 {
		t.Fatal("size is wrong")
	}
	t1 := h.Contains("123")
	t2 := h.Contains("234")
	t3 := h.Contains("345")
	t4 := h.Contains("111")
	t.Log(t1, t2, t3, t4)
	if t1 != true {
		t.Fatal("wrong data")
	}

	if t2 != true {
		t.Fatal("wrong data")
	}

	if t3 != true {
		t.Fatal("wrong data")
	}

	if t4 != false {
		t.Fatal("wrong data")
	}
	h.Remove("123")
	h.Remove("234")
	h.Remove("345")
	if h.Len() != 0 {
		t.Fatal("wrong size")
	}
	h.Add("test")
	h.Clear()
	if h.Len() != 0 {
		t.Fatal("wrong size")
	}

	h.Add("1")
	h.Add("2")
	h.Add("3")
	result := h.ToSlice()
	t.Log(result)
	if len(result) != 3 {
		t.Fatal("wrong data")
	}

}
