package main

import "fmt"

type Node struct {
	value     int
	leftSide  *Node
	rightSide *Node
	parent    *Node
}

func NewNode(value int) (error, *Node) {
	return nil, &Node{
		value: value,
	}
}

type BinaryTree struct {
	root *Node
}

func (bs *BinaryTree) GetRoot() *Node {
	return bs.root
}

func (bs *BinaryTree) Find(key int) *Node {
	current := bs.root
	for current != nil {
		if key < current.value {
			if current.leftSide == nil {
				return current
			}
			current = current.leftSide
		} else if key > current.value {
			if current.rightSide == nil {
				return current
			}
			current = current.rightSide
		} else {
			return current
		}
	}
	return nil
}

func (bs *BinaryTree) Put(value int) {
	if err, newNode := NewNode(value); err == nil {
		if bs.root == nil {
			bs.root = newNode
			return
		} else {
			parent := bs.Find(value)

			if value < parent.value {
				parent.leftSide = newNode
				parent.leftSide.parent = parent
				return
			} else {
				parent.rightSide = newNode
				parent.rightSide.parent = parent
				return
			}
		}
	}
}

func (bs *BinaryTree) PreOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d %s", node.value, " ")
		bs.PreOrder(node.leftSide)
		bs.PreOrder(node.rightSide)
	}
}

func (bs *BinaryTree) InOrder(node *Node) {
	if node != nil {
		bs.InOrder(node.leftSide)
		fmt.Printf("%d %s", node.value, " ")
		bs.InOrder(node.rightSide)
	}
}

func (bs *BinaryTree) PostOrder(node *Node) {
	if node != nil {
		bs.PostOrder(node.leftSide)
		bs.PostOrder(node.rightSide)
		fmt.Printf("%d %s", node.value, " ")
	}
}
