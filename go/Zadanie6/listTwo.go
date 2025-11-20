// listTwo.go
package main

type DoublyNode struct {
	node string
	next *DoublyNode
	prev *DoublyNode
}

type ForwardListTwo struct {
	head *DoublyNode
	tail *DoublyNode
}

func NewForwardListTwo() *ForwardListTwo {
	return &ForwardListTwo{nil, nil}
}

func AddNodeHeadTwo(flist *ForwardListTwo, num string) {
	newNode := &DoublyNode{node: num, next: flist.head, prev: nil}

	if flist.head != nil {
		flist.head.prev = newNode
	}
	flist.head = newNode

	if flist.tail == nil {
		flist.tail = newNode
	}
}

func AddNodeTailTwo(flist *ForwardListTwo, num string) {
	newNode := &DoublyNode{node: num, next: nil, prev: flist.tail}

	if flist.tail != nil {
		flist.tail.next = newNode
	}
	flist.tail = newNode

	if flist.head == nil {
		flist.head = newNode
	}
}

func AddNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
	if target == nil || flist.head == nil {
		return
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
		return
	}

	newNode := &DoublyNode{node: num, next: target.next, prev: target}

	if target.next != nil {
		target.next.prev = newNode
	} else {
		flist.tail = newNode
	}
	target.next = newNode
}

func AddNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
	if target == nil || flist.head == nil {
		return
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
		return
	}

	newNode := &DoublyNode{node: num, next: target, prev: target.prev}

	if target.prev != nil {
		target.prev.next = newNode
	} else {
		flist.head = newNode
	}
	target.prev = newNode
}

func DeleteNodeHeadTwo(flist *ForwardListTwo) {
	if flist.head == nil {
		return
	}

	// Убрана неиспользуемая переменная toDelete
	flist.head = flist.head.next
	if flist.head != nil {
		flist.head.prev = nil
	} else {
		flist.tail = nil
	}
	// Go garbage collector will handle memory
}

func DeleteNodeTailTwo(flist *ForwardListTwo) {
	if flist.tail == nil {
		return
	}

	// Убрана неиспользуемая переменная toDelete
	flist.tail = flist.tail.prev
	if flist.tail != nil {
		flist.tail.next = nil
	} else {
		flist.head = nil
	}
	// Go garbage collector will handle memory
}

func DeleteNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode) {
	if target == nil || flist.head == nil || target.next == nil {
		return
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
		return
	}

	// Убрана неиспользуемая переменная toDelete
	target.next = target.next.next

	if target.next != nil {
		target.next.prev = target
	} else {
		flist.tail = target
	}
	// Go garbage collector will handle memory
}

func DeleteNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode) {
	if target == nil || flist.head == nil || target.prev == nil {
		return
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
		return
	}

	// Убрана неиспользуемая переменная toDelete
	target.prev = target.prev.prev

	if target.prev != nil {
		target.prev.next = target
	} else {
		flist.head = target
	}
	// Go garbage collector will handle memory
}

func DeleteNodeTwo(flist *ForwardListTwo, num string) bool {
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

	// Go garbage collector will handle memory
	return true
}

func FindNodeIndexTwo(flist *ForwardListTwo, num string) bool {
	current := flist.head
	for current != nil {
		if current.node == num {
			return true
		}
		current = current.next
	}
	return false
}

func PrintListTwo(flist *ForwardListTwo) {
	current := flist.head
	print("Список: ")
	for current != nil {
		print(current.node + " ")
		current = current.next
	}
	println()
}

func CountNodesTwo(flist *ForwardListTwo) int {
	count := 0
	current := flist.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func GetNodeByIndexTwo(flist *ForwardListTwo, index int) *DoublyNode {
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

func DeleteNodeIndexTwo(flist *ForwardListTwo, num string) bool {
	return DeleteNodeTwo(flist, num)
}
