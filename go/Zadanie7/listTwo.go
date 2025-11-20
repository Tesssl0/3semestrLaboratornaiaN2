package main

import "fmt"

// DoublyNode представляет узел двусвязного списка
type DoublyNode struct {
	node string
	next *DoublyNode
	prev *DoublyNode
}

// ForwardListTwo представляет двусвязный список
type ForwardListTwo struct {
	head *DoublyNode
	tail *DoublyNode
}

// addNodeHeadTwo добавляет узел в начало списка
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

// addNodeTailTwo добавляет узел в конец списка
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

// addNodeAfterTwo добавляет узел после указанного узла
func addNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
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
		fmt.Println("Ошибка: указанный узел не найден в списке")
		return
	}

	// Вставляем после target
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

// addNodeBeforeTwo добавляет узел перед указанного узла
func addNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode, num string) {
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
		fmt.Println("Ошибка: указанный узел не найден в списке")
		return
	}

	// Вставляем перед target
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

// deleteNodeHeadTwo удаляет узел из начала списка
func deleteNodeHeadTwo(flist *ForwardListTwo) {
	if flist.head == nil {
		return
	}

	toDelete := flist.head
	flist.head = flist.head.next
	if flist.head != nil {
		flist.head.prev = nil
	} else {
		flist.tail = nil
	}
	_ = toDelete // Используем переменную чтобы избежать ошибки
}

// deleteNodeTailTwo удаляет узел из конца списка
func deleteNodeTailTwo(flist *ForwardListTwo) {
	if flist.tail == nil {
		return
	}

	toDelete := flist.tail
	flist.tail = flist.tail.prev
	if flist.tail != nil {
		flist.tail.next = nil
	} else {
		flist.head = nil
	}
	_ = toDelete // Используем переменную чтобы избежать ошибки
}

// deleteNodeAfterTwo удаляет узел после указанного узла
func deleteNodeAfterTwo(flist *ForwardListTwo, target *DoublyNode) {
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

	if !targetFound || target.next == nil {
		return
	}

	toDelete := target.next
	target.next = toDelete.next

	if toDelete.next != nil {
		toDelete.next.prev = target
	} else {
		flist.tail = target
	}
	_ = toDelete // Используем переменную чтобы избежать ошибки
}

// deleteNodeBeforeTwo удаляет узел перед указанного узла
func deleteNodeBeforeTwo(flist *ForwardListTwo, target *DoublyNode) {
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

	if !targetFound || target.prev == nil {
		return
	}

	toDelete := target.prev
	target.prev = toDelete.prev

	if toDelete.prev != nil {
		toDelete.prev.next = target
	} else {
		flist.head = target
	}
	_ = toDelete // Используем переменную чтобы избежать ошибки
}

// deleteNodTwo удаляет узел по значению
func deleteNodTwo(flist *ForwardListTwo, num string) bool {
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

// findNodeIndexTwo ищет узел по значению
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

// printListTwo выводит список
func printListTwo(flist *ForwardListTwo) {
	current := flist.head
	fmt.Print("Список: ")
	for current != nil {
		fmt.Print(current.node + " ")
		current = current.next
	}
	fmt.Println()
}

// countNodesTwo подсчитывает количество узлов в списке
func countNodesTwo(flist *ForwardListTwo) int {
	count := 0
	current := flist.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

// getNodeByIndexTwo возвращает узел по индексу
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

// deleteNodeIndexTwo удаляет узел по значению (аналогично deleteNodTwo)
func deleteNodeIndexTwo(flist *ForwardListTwo, num string) bool {
	return deleteNodTwo(flist, num)
}
