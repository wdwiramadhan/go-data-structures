package stack

import (
	"testing"
)

func TestPush(t *testing.T){
	stack := New()
	if stack.length != 0 && stack.Top != nil {
		t.Fail()
	}
	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	if stack.length == 0 && stack.Top == nil {
		t.Fail()
	}
	if stack.Peek() != 5 {
		t.Fail()
	}
	if stack.Length() != 3 {
		t.Fail()
	}
}

func TestPop(t *testing.T){
	stack := New()
	if stack.length != 0 && stack.Top != nil {
		t.Fail()
	}
	if err := stack.Pop(); err == nil{
		t.Fail()
	}
	stack.Push(5)
	stack.Push(7)
	if err := stack.Pop(); err != nil{
		t.Fail()
	}
}