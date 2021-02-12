package stack

import(
	"fmt"
	"errors"
)

type Stack struct{
	Top *Node
	length int
}

type Node struct{
	Value interface{}
	Prev *Node
}

func New() *Stack{
	return &Stack{nil, 0}
}

func (stack Stack) Peek() interface{}{
	return stack.Top.Value
}

func (stack *Stack) Push(value interface{}) {
	node := &Node{value, stack.Top}
	stack.Top = node
	stack.length++
}

func (stack *Stack) Pop() error{
	if stack.Empty() {
		return errors.New("Stack is empty, push value first")
	}
	stack.Top = stack.Top.Prev
	stack.length--
	return nil
}

func (stack Stack) Empty() bool{
	return stack.length == 0
}

func (stack Stack) Length() int{
	return stack.length
}

func (stack Stack) Display(){
	for stack.Top != nil {
		fmt.Println(stack.Top.Value)
		stack.Top = stack.Top.Prev
	}
}