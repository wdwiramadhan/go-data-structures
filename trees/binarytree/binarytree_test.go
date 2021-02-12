package binarytree

import (
	"testing"
)

func TestInsert(t *testing.T){
	tree := New()
	if tree.Root != nil {
		t.Fail()
	}
	tree.Insert(50)
	if tree.Root.Value != 50 {
		t.Fail()
	}
	tree.Insert(25)
	if tree.Root.Left.Value != 25{
		t.Fail()
	}
	tree.Insert(75)
	if tree.Root.Right.Value != 75{
		t.Fail()
	}
}

func TestSearch(t *testing.T){
	tree := New()
	tree.Insert(50)
	tree.Insert(25)
	if !tree.Search(25) {
		t.Fail()
	}
	if tree.Search(100){
		t.Fail()
	}
}