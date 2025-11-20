package main

import "fmt"

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
	return &ForwardListTwo{nil, nil}
}

func addNodeHeadTwo(flist *ForwardListTwo, num string) {
	newNode := &DoublyNode{
		node: num,
		next: flist.head,
		prev: nil,
	}

	if flist.head != nil {
		flist.head.prev = newNode
	}
	flist.head = newNode

	if flist.tail == nil {
		flist.tail = newNode
	}
}

func addNodeTailTwo(flist *ForwardListTwo, num string) {
	newNode := &DoublyNode{
		node: num,
		next: nil,
		prev: flist.tail,
	}

	if flist.tail != nil {
		flist.tail.next = newNode
	}

	flist.tail = newNode

	if flist.head == nil {
		flist.head = newNode
	}
}

func addNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
	if target == nil || flist.head == nil {
		return
	}

	newNode := &DoublyNode{
		node: num,
		next: target.next,
		prev: target,
	}

	if target.next != nil {
		target.next.prev = newNode
	} else {
		flist.tail = newNode
	}
	target.next = newNode
}

func addNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
	if target == nil || flist.head == nil {
		return
	}

	newNode := &DoublyNode{
		node: num,
		next: target,
		prev: target.prev,
	}

	if target.prev != nil {
		target.prev.next = newNode
	} else {
		flist.head = newNode
	}
	target.prev = newNode
}

func deleteNodeHeadTwo(flist *ForwardListTwo) {
	if flist.head == nil {
		return
	}

	flist.head = flist.head.next
	if flist.head != nil {
		flist.head.prev = nil
	} else {
		flist.tail = nil
	}
	// В Go сборщик мусора автоматически освободит память
}

func deleteNodeTailTwo(flist *ForwardListTwo) {
	if flist.tail == nil {
		return
	}

	flist.tail = flist.tail.prev
	if flist.tail != nil {
		flist.tail.next = nil
	} else {
		flist.head = nil
	}
	// В Go сборщик мусора автоматически освободит память
}

func deleteNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode) {
	if target == nil || target.next == nil {
		return
	}

	target.next = target.next.next
	if target.next != nil {
		target.next.prev = target
	} else {
		flist.tail = target
	}
	// В Go сборщик мусора автоматически освободит память
}

func deleteNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode) {
	if target == nil || target.prev == nil {
		return
	}

	target.prev = target.prev.prev
	if target.prev != nil {
		target.prev.next = target
	} else {
		flist.head = target
	}
	// В Go сборщик мусора автоматически освободит память
}

func deleteNodeTwo(flist *ForwardListTwo, num string) bool {
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

func findNodeIndexTwo(flist *ForwardListTwo, num string) bool {
	current := flist.head
	for current != nil {
		if current.node == num {
			return true
		}
		current = current.next
	}
	return false
}

func printListTwo(flist *ForwardListTwo) {
	current := flist.head
	fmt.Print("Список: ")
	for current != nil {
		fmt.Print(current.node, " ")
		current = current.next
	}
	fmt.Println()
}

func countNodesTwo(flist *ForwardListTwo) int {
	count := 0
	current := flist.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func getNodeByIndexTwo(flist *ForwardListTwo, index int) *DoublyNode {
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

func deleteNodeIndexTwo(flist *ForwardListTwo, num string) bool {
	return deleteNodeTwo(flist, num)
}
