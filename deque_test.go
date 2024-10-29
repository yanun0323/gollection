package gollection

import "testing"

func TestDeque(t *testing.T) {
	d := NewDeque[int]()

	if d.PopFront() != 0 {
		t.Errorf("d.PopFront() = %d, want 0", d.PopFront())
	}

	if d.PopBack() != 0 {
		t.Errorf("d.PopBack() = %d, want 0", d.PopBack())
	}

	d.PushBack(1, 2, 3, 4)
	if d.Len() != 4 {
		t.Errorf("d.Len() = %d, want 4", d.Len())
	}

	for i, s := range d.ToSlice() {
		if s != i+1 {
			t.Errorf("d.ToSlice() = %d, want 0", s)
		}
	}

	if d.PeekBack() != 4 {
		t.Errorf("d.PeekBack() = %d, want 4", d.PeekBack())
	}

	if d.PeekFront() != 1 {
		t.Errorf("d.PeekFront() = %d, want 1", d.PeekFront())
	}

	if d.PopBack() != 4 {
		t.Errorf("d.PopBack() = %d, want 4", d.PopBack())
	}

	if d.PopFront() != 1 {
		t.Errorf("d.PopFront() = %d, want 1", d.PopFront())
	}

	d.PushFront(-1, -2, -3)
	if d.Len() != 5 {
		t.Errorf("d.Len() = %d, want 5", d.Len())
	}

	expected := []int{-3, -2, -1, 2, 3}
	for i, s := range d.ToSlice() {
		if s != expected[i] {
			t.Errorf("d.ToSlice() = %d, want %d", s, expected[i])
		}
	}

	d.Clear()
	if d.Len() != 0 {
		t.Errorf("d.Len() = %d, want 0", d.Len())
	}
}
