package queue

import (
	"fmt"
	"errors"
)

type Node struct{
	Value interface{}
	Next *Node
}

type Queue struct{
	Head *Node
	Tail *Node
	length int
}

func New() *Queue{
	return &Queue{nil,nil, 0}
}

func (queue *Queue) Enqueue(value interface{}) error{
	node := &Node{value, nil}
	if queue.IsEmpty(){
		queue.Head = node
		queue.Tail = node
	} else if queue.Length() == 1 {
		queue.Tail = node
		queue.Head.Next = node
	}else{
		queue.Tail.Next = node
		queue.Tail = node 
	}
	queue.length++
	return nil
}

func (queue *Queue) Dequeue() error{
	if queue.IsEmpty() {
		return errors.New("Queue is empty, enqueue value first")
	}
	if queue.Length() == 1 {
		queue.Head = nil
		queue.Tail = nil
		queue.length = 0
	}else{
		queue.Head = queue.Head.Next
		queue.length--
	}
	return nil
}

func (queue Queue) Length() int {
	return queue.length
}

func (queue Queue) IsEmpty() bool {
	return queue.length == 0 
}

func (queue Queue) Display() {
	for queue.Head != nil {
		fmt.Println(queue.Head.Value)
		queue.Head = queue.Head.Next
	}
}