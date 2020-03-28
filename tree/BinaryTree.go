package main

import "fmt"

type node struct {
	value     int
	leftSide  *node
	rightSide *node
	parent    *node
}

func newNode(value int) (error, *node) {
	return nil, &node{
		value: value,
	}
}

type BinaryTree struct {
	root *node
}

func (bs *BinaryTree) getRoot() *node {
	return bs.root
}

func (bs *BinaryTree) find(key int) *node {
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

func (bs *BinaryTree) put(value int) {
	if err, newNode := newNode(value); err == nil {
		if bs.root == nil {
			bs.root = newNode
			return
		} else {
			parent := bs.find(value)

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

func (bs *BinaryTree) preOrder(node *node) {
	if node != nil {
		fmt.Printf("%d %s", node.value, " ")
		bs.preOrder(node.leftSide)
		bs.preOrder(node.rightSide)
	}
}

func (bs *BinaryTree) inOrder(node *node) {
	if node != nil {
		bs.inOrder(node.leftSide)
		fmt.Printf("%d %s", node.value, " ")
		bs.inOrder(node.rightSide)
	}
}

func (bs *BinaryTree) postOrder(node *node) {
	if node != nil {
		bs.postOrder(node.leftSide)
		bs.postOrder(node.rightSide)
		fmt.Printf("%d %s", node.value, " ")
	}
}
