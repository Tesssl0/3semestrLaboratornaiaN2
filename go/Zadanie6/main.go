// main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sizes := []int{5000, 10000, 20000}

	fmt.Println("=== ЭКСПЕРИМЕНТЫ С ХЕШ-ТАБЛИЦАМИ ===\n")

	hte := &HashTableExperiment{}

	for _, N := range sizes {
		hte.RunMethodComparisonExperiment(N)
	}

	fmt.Println("=== ДЕТАЛЬНЫЕ ЭКСПЕРИМЕНТЫ (МЕТОД ЦЕПОЧЕК) ===\n")
	for _, N := range sizes {
		result := hte.RunSingleExperiment(N, CHAINING)
		hte.PrintResult(result, N)

		if N == 5000 {
			demoTable := NewHashTable(max(1, N/100), CHAINING)
			rg := &RandomGenerator{}
			demoNumbers := rg.GenerateRandomNumbers(100, 0, 1000000)

			for i := 0; i < demoNumbers.Length(); i++ {
				num, _ := strconv.Atoi(demoNumbers.Get(i))
				demoTable.Insert(num)
			}

			fmt.Println("\nДетальная статистика для демо-таблицы:")
			hte.PrintDetailedStatistics(demoTable)
			demoNumbers.Destroy()
		}

		fmt.Printf("\n%s\n\n", strings.Repeat("-", 50))
	}

	hte.RunComparisonExperiment(10000)

	fmt.Println("=== СРАВНЕНИЕ МЕТОДОВ ДЛЯ БОЛЬШОЙ ТАБЛИЦЫ ===")
	hte.RunMethodComparisonExperiment(50000)
}
