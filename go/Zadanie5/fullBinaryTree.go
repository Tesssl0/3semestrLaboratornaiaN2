package main

import (
	"fmt"
	"strconv"
)

// ======== СТРУКТУРА УЗЛА ДЕРЕВА ========
type Node struct {
	data  string
	left  *Node
	right *Node
}

func NewNode(value string) *Node {
	return &Node{
		data:  value,
		left:  nil,
		right: nil,
	}
}

// ======== СТРУКТУРА БИНАРНОГО ДЕРЕВА ========
type FullBinaryTree struct {
	root *Node
}

func NewFullBinaryTree() *FullBinaryTree {
	return &FullBinaryTree{root: nil}
}

// ======== СТРУКТУРА УЗЛА ОЧЕРЕДИ ========
type QueueNode struct {
	treeNode *Node
	next     *QueueNode
}

func NewQueueNode(node *Node) *QueueNode {
	return &QueueNode{
		treeNode: node,
		next:     nil,
	}
}

// ======== ПРОСТАЯ РЕАЛИЗАЦИЯ ОЧЕРЕДИ ========
type SimpleQueue struct {
	front *QueueNode
	rear  *QueueNode
	size  int
}

func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{
		front: nil,
		rear:  nil,
		size:  0,
	}
}

func (q *SimpleQueue) Enqueue(node *Node) {
	newNode := NewQueueNode(node)
	if q.rear == nil {
		q.front = newNode
		q.rear = newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

func (q *SimpleQueue) Dequeue() *Node {
	if q.front == nil {
		return nil
	}

	temp := q.front
	node := temp.treeNode
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	temp = nil // Go автоматически управляет памятью
	q.size--
	return node
}

func (q *SimpleQueue) IsEmpty() bool {
	return q.front == nil
}

func (q *SimpleQueue) Size() int {
	return q.size
}

// ======== ВЫЧИСЛЕНИЕ ВЫСОТЫ ДЕРЕВА ========
func GetTreeHeight(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := GetTreeHeight(node.left)
	rightHeight := GetTreeHeight(node.right)

	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}

// ======== ВСТАВКА В БИНАРНОЕ ДЕРЕВО ПОИСКА ========
func insertBST(node **Node, value string) error {
	if *node == nil {
		*node = NewNode(value)
		return nil
	}

	// Предполагаем, что значения - числа
	newValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	nodeValue, err := strconv.Atoi((*node).data)
	if err != nil {
		return err
	}

	if newValue < nodeValue {
		return insertBST(&(*node).left, value)
	} else if newValue > nodeValue {
		return insertBST(&(*node).right, value)
	}
	// Если значения равны, игнорируем дубликаты
	return nil
}

// Обертка для работы с деревом
func InsertBST(tree *FullBinaryTree, value string) error {
	if tree == nil {
		return fmt.Errorf("дерево не инициализировано")
	}

	// Проверяем, что значение можно преобразовать в число
	if _, err := strconv.Atoi(value); err != nil {
		return fmt.Errorf("ошибка: значение '%s' не является числом", value)
	}

	return insertBST(&tree.root, value)
}

// ======== ПОИСК В ШИРИНУ (BFS) ========
func BFS(tree *FullBinaryTree, value string) *Node {
	if tree == nil || tree.root == nil {
		return nil
	}

	q := NewSimpleQueue()
	q.Enqueue(tree.root)

	for !q.IsEmpty() {
		temp := q.Dequeue()
		if temp.data == value {
			return temp
		}

		if temp.left != nil {
			q.Enqueue(temp.left)
		}
		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}
	return nil
}

// ======== ВЫВОД ВСЕГО ДЕРЕВА В ШИРИНУ ========
func PrintBFS(tree *FullBinaryTree) {
	if tree == nil || tree.root == nil {
		fmt.Println("Дерево пустое")
		return
	}

	q := NewSimpleQueue()
	q.Enqueue(tree.root)

	fmt.Print("Обход в ширину: ")
	for !q.IsEmpty() {
		temp := q.Dequeue()
		fmt.Print(temp.data, " ")

		if temp.left != nil {
			q.Enqueue(temp.left)
		}
		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}
	fmt.Println()
}

// ======== ПОЛУЧЕНИЕ ОБХОДА В ШИРИНУ (BFS) В ВИДЕ СТРОКИ ========
func GetBFSAsString(tree *FullBinaryTree) string {
	if tree == nil || tree.root == nil {
		return ""
	}

	var result string
	q := NewSimpleQueue()
	q.Enqueue(tree.root)

	for !q.IsEmpty() {
		temp := q.Dequeue()
		result += temp.data + " "

		if temp.left != nil {
			q.Enqueue(temp.left)
		}
		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}

	if len(result) > 0 {
		result = result[:len(result)-1] // Убираем последний пробел
	}
	return result
}

// ======== РЕКУРСИВНОЕ УДАЛЕНИЕ ВСЕХ УЗЛОВ ДЕРЕВА ========
func ClearTree(node *Node) {
	if node == nil {
		return
	}
	ClearTree(node.left)
	ClearTree(node.right)
	// В Go нет явного удаления, сборщик мусора сам очистит память
}

// ======== ОЧИСТКА ВСЕГО ДЕРЕВА ========
func ClearFullBinaryTree(tree *FullBinaryTree) {
	if tree == nil {
		return
	}
	ClearTree(tree.root)
	tree.root = nil
}

// ======== ОБХОДЫ ДЕРЕВА ========
func Inorder(node *Node) {
	if node == nil {
		return
	}
	Inorder(node.left)
	fmt.Print(node.data, " ")
	Inorder(node.right)
}

func Preorder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.data, " ")
	Preorder(node.left)
	Preorder(node.right)
}

func Postorder(node *Node) {
	if node == nil {
		return
	}
	Postorder(node.left)
	Postorder(node.right)
	fmt.Print(node.data, " ")
}

// ======== ПРОВЕРКА НА ПОЛНОЕ БИНАРНОЕ ДЕРЕВО ========
func IsCompleteBinaryTree(root *Node) bool {
	if root == nil {
		return true
	}

	q := NewSimpleQueue()
	q.Enqueue(root)
	foundNull := false

	for !q.IsEmpty() {
		temp := q.Dequeue()

		if temp.left != nil {
			if foundNull {
				return false
			}
			q.Enqueue(temp.left)
		} else {
			foundNull = true
		}

		if temp.right != nil {
			if foundNull {
				return false
			}
			q.Enqueue(temp.right)
		} else {
			foundNull = true
		}
	}
	return true
}

// ======== ПРОВЕРКА НА СТРОГО ПОЛНОЕ ДЕРЕВО ========
func IsFullBinaryTree(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return true
	}
	if root.left != nil && root.right != nil {
		return IsFullBinaryTree(root.left) && IsFullBinaryTree(root.right)
	}
	return false
}

// ======== ПОЛУЧЕНИЕ ОБХОДА INORDER В ВИДЕ СТРОКИ ========
func getInorderHelper(node *Node, result *string) {
	if node == nil {
		return
	}
	getInorderHelper(node.left, result)
	*result += node.data + " "
	getInorderHelper(node.right, result)
}

func GetInorderAsString(node *Node) string {
	var result string
	getInorderHelper(node, &result)
	if len(result) > 0 {
		result = result[:len(result)-1] // Убираем последний пробел
	}
	return result
}

// ======== АНАЛИЗ ДЕРЕВА (ТИП, ВЫСОТА, УЗЛЫ) ========
func CheckTreeType(tree *FullBinaryTree) {
	if tree == nil || tree.root == nil {
		fmt.Println("The tree is empty")
		return
	}

	complete := IsCompleteBinaryTree(tree.root)
	full := IsFullBinaryTree(tree.root)

	fmt.Println("=== TREE ANALYSIS ===")

	// Подсчет количества узлов
	count := 0
	q := NewSimpleQueue()
	if tree.root != nil {
		q.Enqueue(tree.root)
	}
	for !q.IsEmpty() {
		temp := q.Dequeue()
		count++
		if temp.left != nil {
			q.Enqueue(temp.left)
		}
		if temp.right != nil {
			q.Enqueue(temp.right)
		}
	}
	fmt.Printf("Number of nodes: %d\n", count)

	// Вычисление высоты дерева (по уровням)
	height := 0
	if tree.root != nil {
		q2 := NewSimpleQueue()
		q2.Enqueue(tree.root)

		for !q2.IsEmpty() {
			levelSize := q2.size
			height++
			for i := 0; i < levelSize; i++ {
				temp := q2.Dequeue()
				if temp.left != nil {
					q2.Enqueue(temp.left)
				}
				if temp.right != nil {
					q2.Enqueue(temp.right)
				}
			}
		}
	}

	fmt.Printf("Tree height: %d\n", height)

	fmt.Print("Tree type: ")
	if full {
		fmt.Println("Full Binary Tree")
	} else {
		fmt.Println("Binary Tree")
		fmt.Println("- Does not meet the criteria for a full binary tree")
	}

	maxNodes := (1 << height) - 1
	fmt.Printf("Maximum possible nodes for height %d: %d\n", height, maxNodes)
	fmt.Printf("Actual number of nodes: %d\n", count)

	if complete && count == maxNodes {
		fmt.Println("The tree is a PERFECT binary tree!")
	}
}

// ======== ЗМЕЙКА ========
func PrintZigzag(tree *FullBinaryTree) {
	if tree == nil || tree.root == nil {
		fmt.Println("Дерево пустое")
		return
	}

	q := NewSimpleQueue()
	q.Enqueue(tree.root)
	leftToRight := true // Направление вывода

	fmt.Print("Обход змейкой: ")

	for !q.IsEmpty() {
		levelSize := q.size

		// Временный массив для хранения узлов текущего уровня
		levelNodes := make([]string, levelSize)

		// Обрабатываем все узлы текущего уровня
		for i := 0; i < levelSize; i++ {
			temp := q.Dequeue()

			// Сохраняем узел в массив в зависимости от направления
			if leftToRight {
				levelNodes[i] = temp.data
			} else {
				levelNodes[levelSize-1-i] = temp.data
			}

			// Добавляем дочерние узлы в очередь
			if temp.left != nil {
				q.Enqueue(temp.left)
			}
			if temp.right != nil {
				q.Enqueue(temp.right)
			}
		}

		// Выводим узлы текущего уровня
		for i := 0; i < levelSize; i++ {
			fmt.Print(levelNodes[i])
			if i < levelSize-1 {
				fmt.Print(" - ")
			}
		}

		// Меняем направление для следующего уровня
		leftToRight = !leftToRight

		// Добавляем разделитель между уровнями (кроме последнего)
		if !q.IsEmpty() {
			fmt.Print(" - ")
		}
	}
	fmt.Println()
}
