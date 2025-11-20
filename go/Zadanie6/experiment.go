// experiment.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ExperimentResult struct {
	MinLength  int
	MaxLength  int
	AvgLength  float64
	GenTime    int64
	InsertTime int64
	TableSize  int
	LoadFactor float64
	Method     CollisionResolution
}

type HashTableExperiment struct{}

func (hte *HashTableExperiment) RunSingleExperiment(N int, method CollisionResolution) ExperimentResult {
	result := ExperimentResult{}
	result.Method = method

	if method == CHAINING {
		result.TableSize = max(1, N/100)
	} else {
		result.TableSize = max(1, N*2)
	}

	result.LoadFactor = float64(N) / float64(result.TableSize)

	timer := &Timer{}
	rg := &RandomGenerator{}

	timer.Start()
	randomNumbers := rg.GenerateRandomNumbers(N, 0, 1000000)
	result.GenTime = timer.ElapsedMillis()

	hashTable := NewHashTable(result.TableSize, method)

	timer.Start()
	for i := 0; i < randomNumbers.Length(); i++ {
		num, _ := strconv.Atoi(randomNumbers.Get(i))
		hashTable.Insert(num)
	}
	result.InsertTime = timer.ElapsedMillis()

	minLen, maxLen, avgLen := hashTable.GetStatistics()
	result.MinLength = minLen
	result.MaxLength = maxLen
	result.AvgLength = avgLen

	randomNumbers.Destroy()

	return result
}

func (hte *HashTableExperiment) RunComparisonExperiment(N int) {
	fmt.Printf("=== Сравнительный эксперимент для N = %d ===\n", N)

	rg := &RandomGenerator{}
	randomNumbers := rg.GenerateRandomNumbers(N, 0, 1000000)

	tableSizes := []int{N / 200, N / 100, N / 50, N / 20}

	for _, tableSize := range tableSizes {
		if tableSize < 1 {
			tableSize = 1
		}

		hashTable := NewHashTable(tableSize, CHAINING)

		for i := 0; i < randomNumbers.Length(); i++ {
			num, _ := strconv.Atoi(randomNumbers.Get(i))
			hashTable.Insert(num)
		}

		minLen, maxLen, avgLen := hashTable.GetStatistics()

		fmt.Printf("Размер таблицы: %6d | Заполнение: %8.2f | Цепочки: мин=%3d, макс=%3d, средняя=%6.2f\n",
			tableSize, float64(N)/float64(tableSize), minLen, maxLen, avgLen)
	}

	randomNumbers.Destroy()
	fmt.Printf("\n%s\n\n", strings.Repeat("=", 50))
}

func (hte *HashTableExperiment) RunMethodComparisonExperiment(N int) {
	fmt.Printf("=== Сравнение методов разрешения коллизий для N = %d ===\n", N)

	methods := []CollisionResolution{CHAINING, LINEAR_PROBING, QUADRATIC_PROBING}
	methodNames := []string{"Метод цепочек", "Линейное пробирование", "Квадратичное пробирование"}

	for i := 0; i < 3; i++ {
		result := hte.RunSingleExperiment(N, methods[i])

		fmt.Printf("%s:\n", methodNames[i])
		fmt.Printf("  Время вставки: %d мс\n", result.InsertTime)
		fmt.Printf("  Цепочки: мин=%d, макс=%d, средняя=%.2f\n",
			result.MinLength, result.MaxLength, result.AvgLength)
		fmt.Printf("  Коэф. заполнения: %.2f\n\n", result.LoadFactor)
	}

	fmt.Printf("%s\n\n", strings.Repeat("=", 50))
}

func (hte *HashTableExperiment) PrintResult(result ExperimentResult, N int) {
	methodName := ""
	switch result.Method {
	case CHAINING:
		methodName = "Метод цепочек"
	case LINEAR_PROBING:
		methodName = "Линейное пробирование"
	case QUADRATIC_PROBING:
		methodName = "Квадратичное пробирование"
	}

	fmt.Printf("=== Эксперимент для N = %d (%s) ===\n", N, methodName)
	fmt.Printf("Размер хеш-таблицы: %d\n", result.TableSize)
	fmt.Printf("Коэффициент заполнения: %.2f\n", result.LoadFactor)
	fmt.Printf("Время генерации чисел: %d мс\n", result.GenTime)
	fmt.Printf("Время вставки: %d мс\n", result.InsertTime)
	fmt.Printf("\nРезультаты:\n")
	fmt.Printf("Самая короткая цепочка: %d\n", result.MinLength)
	fmt.Printf("Самая длинная цепочка: %d\n", result.MaxLength)
	fmt.Printf("Средняя длина непустой цепочки: %.2f\n", result.AvgLength)
}

func (hte *HashTableExperiment) PrintDetailedStatistics(hashTable *HashTable) {
	hashTable.PrintDetailedStatistics()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
