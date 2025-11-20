package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tree := NewFullBinaryTree()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите последовательность целых чисел (для завершения введите 'q'):")

	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == "q" || input == "Q" {
			break
		}

		// Проверяем, что введено число
		if _, err := strconv.Atoi(input); err != nil {
			fmt.Printf("Ошибка: '%s' не является числом\n", input)
			break
		}

		// Если числа нет в дереве, добавляем его
		if BFS(tree, input) == nil {
			if err := InsertBST(tree, input); err != nil {
				fmt.Println(err)
				break
			}
		}
	}

	// Вычисляем высоту дерева с помощью функции из библиотеки
	height := GetTreeHeight(tree.root)
	fmt.Printf("Высота получившегося дерева: %d\n", height)

	// Очищаем память
	ClearFullBinaryTree(tree)
	fmt.Println("\nДерево очищено")
}
