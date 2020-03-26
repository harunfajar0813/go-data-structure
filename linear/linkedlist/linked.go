package main

import "fmt"

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

type linked struct {
	head   *node
	tail   *node
	length int
}

func (l *linked) circularConnection() {
	l.head.previous = l.tail
	l.head.previous.next = l.head
}

func (l *linked) insert(value int) {
	if err, newNode := newNode(value); err == nil {
		if l.tail == nil {
			l.tail = newNode
			l.head = l.tail
		} else {
			newNode.next = l.head
			newNode.next.previous = newNode
			l.head = l.head.previous
		}
		l.circularConnection()
		l.length++
	}
}

func (l *linked) append(value int) {
	if err, newNode := newNode(value); err == nil {
		if l.tail == nil {
			l.tail = newNode
			l.head = l.tail
		} else {
			newNode.previous = l.tail
			newNode.previous.next = newNode
			l.tail = l.tail.next
		}
		l.circularConnection()
		l.length++
	}
}

func (l *linked) isEmpty() bool {
	return l.tail == nil
}

func (l *linked) isOnlyOneNode() bool {
	return l.tail == l.head
}

func (l *linked) deleteHead() error {
	if l.isEmpty() {
		return fmt.Errorf("list is empty")
	} else {
		if l.isOnlyOneNode() {
			l.head, l.tail = nil, nil
		} else {
			l.head = l.head.next
			l.head.previous = nil
			l.circularConnection()
		}
	}
	l.length--
	return nil
}

func (l *linked) deleteTail() error {
	if l.isEmpty() {
		return fmt.Errorf("list is empty")
	} else {
		if l.isOnlyOneNode() {
			l.tail, l.head = nil, nil
		} else {
			l.tail = l.tail.previous
			l.tail.next = nil
			l.circularConnection()
		}
	}
	l.length--
	return nil
}

func (l *linked) show() {
	if l.isEmpty() {
		fmt.Errorf("list is empty")
		return
	}
	current := l.head
	for index := 0; index < l.length; index++ {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
}

func (l *linked) showReverse() {
	if l.isEmpty() {
		fmt.Errorf("list is empty")
		return
	}
	current := l.tail
	for index := 0; index < l.length; index++ {
		fmt.Print(current.value, " -> ")
		current = current.next
	}
}
