package binarytree

type BinaryNode struct{
	Value int32
	Left *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct{
	Root *BinaryNode
}

func New() *BinaryTree{
	return &BinaryTree{nil}
}

func (t *BinaryTree) Insert(value int32) {
	if t.Root == nil {
		t.Root = &BinaryNode{value, nil, nil}
	}else{
		t.Root.InsertNode(value)
	}
}

func (n *BinaryNode) InsertNode(value int32){
	if n == nil {
		return
	}else if value <= n.Value {
		if n.Left == nil {
			n.Left = &BinaryNode{value, nil, nil}
		}else{
			n.Left.InsertNode(value)
		}
	}else{
		if n.Right == nil {
			n.Right = &BinaryNode{value, nil, nil}
		}else{
			n.Right.InsertNode(value)
		}
	}
}

func (t BinaryTree) Search(key int32) bool{
	current := t.Root
	for current != nil {
		if current.Value == key {
			return true
		}
		if key <= current.Value {
			current = current.Left
		}else{
			current = current.Right
		}
	}
	return false
}