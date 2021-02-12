package singlelinkedlist

import (
	"testing"
)

func TestInsertFirst(t *testing.T){
	list := New()
	if list.Head != nil && list.length != 0 {
		t.Fail()
	}
	list.InsertFirst(3)
	if list.Head.Value != 3 {
		t.Fail()
	}
	if list.length != 1 {
		t.Fail()
	} 
}

func TestDeleteFirst(t *testing.T){
	list := New()
	if err := list.DeleteFirst(); err == nil {
		t.Fail()
	}
	list.InsertFirst(5)
	list.InsertFirst(7)
	if err := list.DeleteFirst(); err != nil {
		t.Fail()
	}
	if list.Length() != 1 {
		t.Fail()
	}
}