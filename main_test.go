package main

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	head := NewNode(-1, nil, MaxLevel)

	if head.key != -1 {
		t.Errorf("Actual %v, expected %v", head.key, -1)
	}
	if head.value != nil {
		t.Errorf("Actual %v, expected %v", head.value, nil)
	}
	if len(head.level) != MaxLevel {
		t.Errorf("Actual %v, expected %v", len(head.level), MaxLevel)
	}
}

func TestNewSkipList(t *testing.T) {
	sl := NewSkipList()
	if sl.level != 1 {
		t.Errorf("Actual %v, expected %v", sl.level, -1)
	}

	if sl.head.key != -1 {
		t.Errorf("Actual %v, expected %v", sl.head, -1)
	}

	if sl.head.value != nil {
		t.Errorf("Actual %v, expected %v", sl.head, -1)
	}

	if len(sl.head.level) != MaxLevel {
		t.Errorf("Actual %v, expected %v", sl.head, -1)
	}
}

func TestSkipList_Insert(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(3, "three")
	sl.Insert(6, "six")
	sl.Insert(7, "seven")
	sl.Insert(9, "nine")

	if sl.Search(3) == nil {
		t.Errorf("Actual %v, expected %v", nil, 3)
	}
	if sl.Search(9) == nil {
		t.Errorf("Actual %v, expected %v", nil, 9)
	}
}

func TestSkipList_Delete(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(3, "three")
	sl.Insert(6, "six")

	sl.Delete(3)

	if sl.Search(3) != nil {
		t.Errorf("Actual %v, expected %v", 3, nil)
	}

	if len(sl.head.level) != MaxLevel {
		t.Errorf("Actual %v, expected %v", len(sl.head.level), 1)
	}

}

func TestSkipList_Search(t *testing.T) {
	sl := NewSkipList()
	sl.Insert(3, "three")
	sl.Insert(6, "six")

	if sl.Search(3) == nil {
		t.Errorf("Actual %v, expected %v", nil, 3)
	}
}
