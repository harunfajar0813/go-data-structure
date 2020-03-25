package main

type node struct {
	value    int
	next     *node
	previous *node
}

func newNode(value int) (error, *node) {
	return nil, &node{
		value: value,
	}
}
