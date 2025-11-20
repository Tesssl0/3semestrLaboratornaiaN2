// listTwo.go
package main

import (
	"fmt"
)

type DoublyNode struct {
	node string
	next *DoublyNode
	prev *DoublyNode
}

type ForwardListTwo struct {
	head *DoublyNode
	tail *DoublyNode
}

func newForwardListTwo() *ForwardListTwo {
	return &ForwardListTwo{}
}

func (flist *ForwardListTwo) addNodeHead(num string) {
	newNode := &DoublyNode{node: num, next: flist.head}

	if flist.head != nil {
		flist.head.prev = newNode
	}
	flist.head = newNode

	if flist.tail == nil {
		flist.tail = newNode
	}
}

func (flist *ForwardListTwo) addNodeTail(num string) {
	newNode := &DoublyNode{node: num, prev: flist.tail}

	if flist.tail != nil {
		flist.tail.next = newNode
	}
	flist.tail = newNode

	if flist.head == nil {
		flist.head = newNode
	}
}

func (flist *ForwardListTwo) addNodeAfter(target *DoublyNode, num string) bool {
	if target == nil || flist.head == nil {
		return false
	}

	// Проверяем, что target находится в списке
	current := flist.head
	targetFound := false
	for current != nil {
		if current == target {
			targetFound = true
			break
		}
		current = current.next
	}

	if !targetFound {
		return false
	}

	newNode := &DoublyNode{node: num, next: target.next, prev: target}

	if target.next != nil {
		target.next.prev = newNode
	} else {
		flist.tail = newNode
	}
	target.next = newNode

	return true
}

func (flist *ForwardListTwo) addNodeBefore(target *DoublyNode, num string) bool {
	if target == nil || flist.head == nil {
		return false
	}

	// Проверяем, что target находится в списке
	current := flist.head
	targetFound := false
	for current != nil {
		if current == target {
			targetFound = true
			break
		}
		current = current.next
	}

	if !targetFound {
		return false
	}

	newNode := &DoublyNode{node: num, next: target, prev: target.prev}

	if target.prev != nil {
		target.prev.next = newNode
	} else {
		flist.head = newNode
	}
	target.prev = newNode

	return true
}

func (flist *ForwardListTwo) deleteNodeHead() {
	if flist.head == nil {
		return
	}

	// Просто перенаправляем указатели, сборщик мусора Go удалит узел автоматически
	flist.head = flist.head.next
	if flist.head != nil {
		flist.head.prev = nil
	} else {
		flist.tail = nil
	}
}

func (flist *ForwardListTwo) deleteNodeTail() {
	if flist.tail == nil {
		return
	}

	// Просто перенаправляем указатели, сборщик мусора Go удалит узел автоматически
	flist.tail = flist.tail.prev
	if flist.tail != nil {
		flist.tail.next = nil
	} else {
		flist.head = nil
	}
}

func (flist *ForwardListTwo) deleteNodeAfter(target *DoublyNode) bool {
	if target == nil || flist.head == nil || target.next == nil {
		return false
	}

	// Проверяем, что target находится в списке
	current := flist.head
	targetFound := false
	for current != nil {
		if current == target {
			targetFound = true
			break
		}
		current = current.next
	}

	if !targetFound {
		return false
	}

	// Перенаправляем указатели
	nodeToDelete := target.next
	target.next = nodeToDelete.next

	if nodeToDelete.next != nil {
		nodeToDelete.next.prev = target
	} else {
		flist.tail = target
	}

	return true
}

func (flist *ForwardListTwo) deleteNodeBefore(target *DoublyNode) bool {
	if target == nil || flist.head == nil || target.prev == nil {
		return false
	}

	// Проверяем, что target находится в списке
	current := flist.head
	targetFound := false
	for current != nil {
		if current == target {
			targetFound = true
			break
		}
		current = current.next
	}

	if !targetFound {
		return false
	}

	// Перенаправляем указатели
	nodeToDelete := target.prev
	target.prev = nodeToDelete.prev

	if nodeToDelete.prev != nil {
		nodeToDelete.prev.next = target
	} else {
		flist.head = target
	}

	return true
}

func (flist *ForwardListTwo) deleteNode(num string) bool {
	if flist.head == nil {
		return false
	}

	current := flist.head
	for current != nil && current.node != num {
		current = current.next
	}

	if current == nil {
		return false
	}

	if current.prev != nil {
		current.prev.next = current.next
	} else {
		flist.head = current.next
	}

	if current.next != nil {
		current.next.prev = current.prev
	} else {
		flist.tail = current.prev
	}

	return true
}

func (flist *ForwardListTwo) findNodeIndex(num string) bool {
	current := flist.head
	for current != nil {
		if current.node == num {
			return true
		}
		current = current.next
	}
	return false
}

func (flist *ForwardListTwo) printList() {
	current := flist.head
	fmt.Print("Список: ")
	for current != nil {
		fmt.Print(current.node, " ")
		current = current.next
	}
	fmt.Println()
}

func (flist *ForwardListTwo) countNodes() int {
	count := 0
	current := flist.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (flist *ForwardListTwo) getNodeByIndex(index int) *DoublyNode {
	if index < 0 {
		return nil
	}

	current := flist.head
	currentIndex := 0

	for current != nil {
		if currentIndex == index {
			return current
		}
		current = current.next
		currentIndex++
	}

	return nil
}
