// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Ошибка: не удалось открыть файл %s", filename)
	}
	defer file.Close()

	fmt.Printf("Содержимое файла %s:\n", filename)
	fmt.Println("----------------------------------------")

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			fmt.Printf("%d: %s\n", lineNumber, line)
			lineNumber++
		}
	}

	if lineNumber == 1 {
		fmt.Println("Файл пуст")
	}
	fmt.Println("----------------------------------------")

	return scanner.Err()
}

func printUsage() {
	fmt.Println("Использование: ./program --file <файл> --query <операция>")
	fmt.Println("Доступные команды:")
	fmt.Println("  SETADD <число>  - добавить элемент в множество")
	fmt.Println("  SETDEL <число>  - удалить элемент из множества")
	fmt.Println("  SET_AT <число>  - проверить наличие элемента")
	fmt.Println("  SHOW_FILE       - вывести содержимое файла")
}

func main() {
	if len(os.Args) != 5 {
		printUsage()
		os.Exit(1)
	}

	filename := ""
	query := ""

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "--file" && i+1 < len(os.Args) {
			filename = os.Args[i+1]
			i++
		} else if arg == "--query" && i+1 < len(os.Args) {
			query = os.Args[i+1]
			i++
		}
	}

	if filename == "" || query == "" {
		printUsage()
		os.Exit(1)
	}

	mySet := NewSet(100)
	err := mySet.LoadFromFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	parts := strings.Fields(query)
	if len(parts) == 0 {
		fmt.Println("Ошибка: пустой запрос")
		os.Exit(1)
	}

	cmd := parts[0]

	if cmd == "SHOW_FILE" {
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if len(parts) < 2 {
		fmt.Println("Ошибка: не указано значение для операции")
		os.Exit(1)
	}

	valueStr := parts[1]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Ошибка: некорректное значение числа")
		os.Exit(1)
	}

	switch cmd {
	case "SETADD":
		if mySet.Add(value) {
			fmt.Printf("Элемент %d добавлен\n", value)
			err := mySet.SaveToFile(filename)
			if err != nil {
				fmt.Println("Ошибка при сохранении файла:", err)
			}
		} else {
			fmt.Printf("Элемент %d уже существует\n", value)
		}
	case "SETDEL":
		if mySet.Remove(value) {
			fmt.Printf("Элемент %d удалён\n", value)
			err := mySet.SaveToFile(filename)
			if err != nil {
				fmt.Println("Ошибка при сохранении файла:", err)
			}
		} else {
			fmt.Printf("Элемент %d не найден\n", value)
		}
	case "SET_AT":
		if mySet.Contains(value) {
			fmt.Printf("Элемент %d присутствует в множестве\n", value)
		} else {
			fmt.Printf("Элемент %d отсутствует в множестве\n", value)
		}
	default:
		fmt.Printf("Неизвестная операция: %s\n", cmd)
		printUsage()
		os.Exit(1)
	}
}
