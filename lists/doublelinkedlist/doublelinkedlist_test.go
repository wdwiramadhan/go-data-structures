package doublelinkedlist

import (
	"testing"
)

func TestInsertFirst(t *testing.T){
	list := New()
	if list.Head != nil && list.Tail != nil && list.length != 0 {
		t.Fail()
	}
	list.InsertFirst(3)
	if list.Head.Data != 3 {
		t.Fail()
	}
	if list.Length() != 1 {
		t.Fail()
	}
	list.InsertFirst(5)
	if list.Head.Data != 5 {
		t.Fail()
	}
	if list.Head.Next.Data != 3 {
		t.Fail()
	}
	if list.Tail.Data != 3 {
		t.Fail()
	}
	if list.Length() != 2 {
		t.Fail()
	}
}

func TestInsertLast(t *testing.T){
	list := New()
	list.InsertLast(7)
	if list.Tail.Data != 7 {
		t.Fail()
	}
	if list.Length() != 1 {
		t.Fail()
	}
	list.InsertFirst(5)
	list.InsertFirst(3)
	if list.Head.Data != 3 {
		t.Fail()
	}
}

func TestDeleteFirst(t *testing.T){
	list := New()
	if err := list.DeleteFirst(); err == nil {
		t.Fail()
	}
	list.InsertFirst(3)
	list.DeleteFirst()
	if list.IsEmpty() != true {
		t.Fail()
	}
	list.InsertFirst(3)
	list.InsertLast(5)
	list.InsertLast(7)
	list.DeleteFirst()
	if list.Head.Data != 5 {
		t.Fail()
	}
	list.DeleteLast()
	if list.Head.Data != 5 {
		t.Fail()
	}
}