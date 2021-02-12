package singlelinkedlist

import (
	"fmt"
	"errors"
)

type Node struct{
	Value interface{}
	Next *Node
}

type List struct{
	Head *Node
	length int
}

func New() *List{
	return &List{nil, 0}
}

func (list *List)  InsertFirst(value interface{}) error {
	node := &Node{value, nil}
	if list.IsEmpty() {
		list.Head = node
		list.length++
		return nil
	}
	node.Next = list.Head
	list.Head = node
	list.length++
	return nil
}

func (list *List) InsertLast(value interface{}) error {
	node := &Node{value, nil}
	if list.IsEmpty() {
		list.Head = node
		list.length++
		return nil
	}

	temp := list.Head
	for temp.Next != nil{
		temp = temp.Next
	}
	temp.Next = node
	return nil
}

func (list *List) DeleteFirst() error {
	if list.IsEmpty() {
		return errors.New("List is empty, insert value first")
	}
	list.Head = list.Head.Next
	list.length--
	return nil
}

func (list *List) DeleteLast() error {
	if list.IsEmpty() {
		return errors.New("List is empty, insert value first")
	}
	temp := list.Head
	for temp.Next.Next != nil{
		temp = temp.Next
	}
	temp.Next = nil
	return nil
}

func (list List) Length() int {
	return list.length
}

func (list List) IsEmpty() bool {
	return list.length == 0
}

func (list List) Display() {
	for list.Head != nil {
		fmt.Println(list.Head.Value)
		list.Head = list.Head.Next
	}
}