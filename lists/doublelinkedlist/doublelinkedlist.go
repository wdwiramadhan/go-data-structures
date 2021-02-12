package doublelinkedlist

import (
	"fmt"
	"errors"
)

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

type List struct {
	Head *Node
	Tail *Node
	length int
}

func New() *List {
	return &List{nil, nil, 0}
}

func (list *List) InsertFirst(data interface{}) error{
	node := &Node{data, nil, nil}
	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		list.length++
		return nil
	}
	list.Head.Prev = node
	node.Next = list.Head
	list.Head = node
	list.length++
	return nil
}

func (list *List) InsertLast(data interface{}) error{
	node := &Node{data, nil, nil}
	if list.IsEmpty() {
		list.Head = node
		list.Tail = node
		list.length++
		return nil
	}
	if list.Length() == 1 {
		list.Head.Next = node
		node.Prev = list.Head
		list.length++
		list.Tail = node
		return nil
	}
	node.Prev = list.Tail
	list.Tail.Next = node
	list.Tail = node
	list.length++
	return nil
}

func (list *List) DeleteFirst() error {
	if list.IsEmpty() {
		return errors.New("List is empty, insert data first")
	}
	if list.Length() == 1 {
		list.Head = nil
		list.Tail = nil
		list.length = 0
		return nil
	}
	list.Head = list.Head.Next
	list.Head.Prev = nil
	list.length--
	return nil
}

func (list *List) DeleteLast() error {
	if list.IsEmpty() {
		return errors.New("List is empty, insert data first")
	}
	if list.Length() == 1 {
		list.Head = nil
		list.Tail = nil
		list.length = 0
		return nil
	}
	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	list.length--
	return nil
}

func (list List) IsEmpty() bool {
	return list.length == 0
}

func (list List) Length() int{
	return list.length
}

func (list List) Display(){
	for list.Head != nil {
		fmt.Println(list.Head.Data)
		list.Head = list.Head.Next
	}
}