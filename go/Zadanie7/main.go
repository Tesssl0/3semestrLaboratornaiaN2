package main

import "fmt"

func runTestCases() {
	fmt.Println("=== Тест 1 ===")
	fmt.Println("cap = 2, Q = 2")
	fmt.Println("Queries = SET 1 2 GET 1")

	cache1 := NewLRUCache(2)
	cache1.set(1, 2)
	cache1.printCache()
	result1 := cache1.get(1)
	fmt.Println("GET 1 =", result1)
	fmt.Println("Ожидается: 2")

	fmt.Println("\n=== Тест 2 ===")
	fmt.Println("cap = 2, Q = 8")
	fmt.Println("Queries = SET 1 2 SET 2 3 SET 1 5 SET 4 5 SET 6 7 GET 4 SET 1 2 GET 3")

	cache2 := NewLRUCache(2)

	// Храним результаты GET запросов в массиве
	var results []int

	cache2.set(1, 2)
	cache2.printCache()
	cache2.set(2, 3)
	cache2.printCache()
	cache2.set(1, 5)
	cache2.printCache()
	cache2.set(4, 5)
	cache2.printCache()
	cache2.set(6, 7)
	cache2.printCache()

	results = append(results, cache2.get(4))
	fmt.Println("GET 4 =", results[len(results)-1])

	cache2.set(1, 2)
	cache2.printCache()

	results = append(results, cache2.get(3))
	fmt.Println("GET 3 =", results[len(results)-1])

	fmt.Print("Результат: ")
	for i, result := range results {
		fmt.Print(result)
		if i < len(results)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
	fmt.Println("Ожидается: 5 -1")
}

func main() {
	var choice int

	for {
		fmt.Println("\n=== LRU Cache Menu ===")
		fmt.Println("1. Запустить тестовые примеры")
		fmt.Println("2. Ввести свои запросы")
		fmt.Println("3. Выход")
		fmt.Print("Выберите опцию: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			runTestCases()
		case 2:
			processLRUQueries()
		case 3:
			fmt.Println("Выход...")
			return
		default:
			fmt.Println("Неверный выбор!")
		}
	}
}
