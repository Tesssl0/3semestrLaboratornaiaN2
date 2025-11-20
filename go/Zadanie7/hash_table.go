package main

import (
	"fmt"
	"math"
	"strconv"
)

// ChainingHashTable представляет хеш-таблицу с методом цепочек
type ChainingHashTable struct {
	table []*ForwardListTwo
	size  int
}

// hash вычисляет хеш-код для ключа
func (ht *ChainingHashTable) hash(key int) int {
	return int(math.Abs(float64(key))) % ht.size
}

// NewChainingHashTable создает новую хеш-таблицу
func NewChainingHashTable(tableSize int) *ChainingHashTable {
	ht := &ChainingHashTable{
		table: make([]*ForwardListTwo, tableSize),
		size:  tableSize,
	}

	// Инициализируем каждую ячейку таблицы пустым списком
	for i := 0; i < tableSize; i++ {
		ht.table[i] = &ForwardListTwo{
			head: nil,
			tail: nil,
		}
	}

	return ht
}

// insert вставляет элемент в хеш-таблицу
func (ht *ChainingHashTable) insert(key int) {
	index := ht.hash(key)
	list := ht.table[index]
	if list != nil {
		addNodeTailTwo(list, strconv.Itoa(key))
	}
}

// search ищет элемент в хеш-таблице
func (ht *ChainingHashTable) search(key int) bool {
	index := ht.hash(key)
	list := ht.table[index]
	if list == nil {
		return false
	}
	return findNodeIndexTwo(list, strconv.Itoa(key))
}

// getStatistics возвращает статистику по хеш-таблице
func (ht *ChainingHashTable) getStatistics() (minLength, maxLength int, avgLength float64) {
	minLength = math.MaxInt32
	maxLength = 0
	totalElements := 0
	nonEmptyBuckets := 0

	// Проходим по всем ячейкам таблицы
	for i := 0; i < ht.size; i++ {
		list := ht.table[i]
		if list != nil {
			listSize := countNodesTwo(list)
			if listSize > 0 {
				// Обновляем минимальную и максимальную длину
				if listSize < minLength {
					minLength = listSize
				}
				if listSize > maxLength {
					maxLength = listSize
				}
				totalElements += listSize
				nonEmptyBuckets++
			}
		}
	}

	// Корректируем значения если нет непустых ячеек
	if minLength == math.MaxInt32 {
		minLength = 0
	}

	// Вычисляем среднюю длину цепочек
	if nonEmptyBuckets > 0 {
		avgLength = float64(totalElements) / float64(nonEmptyBuckets)
	} else {
		avgLength = 0
	}

	return minLength, maxLength, avgLength
}

// remove удаляет элемент из хеш-таблицы
func (ht *ChainingHashTable) remove(key int) {
	index := ht.hash(key)
	list := ht.table[index]
	if list != nil {
		deleteNodeIndexTwo(list, strconv.Itoa(key))
	}
}

// printDetailedStatistics выводит детальную статистику
func (ht *ChainingHashTable) printDetailedStatistics() {
	totalElements := 0
	emptyBuckets := 0
	maxChain := 0
	minChain := math.MaxInt32

	fmt.Println("Распределение цепочек:")

	// Анализируем каждую ячейку таблицы
	for i := 0; i < ht.size; i++ {
		list := ht.table[i]
		chainLength := 0
		if list != nil {
			chainLength = countNodesTwo(list)
		}

		// Собираем статистику
		if chainLength == 0 {
			emptyBuckets++
		}
		if chainLength > maxChain {
			maxChain = chainLength
		}
		if chainLength < minChain && chainLength > 0 {
			minChain = chainLength
		}

		totalElements += chainLength

		// Показываем первые 10 цепочек для примера
		if i < 10 {
			fmt.Printf("Ячейка %d: %d элементов\n", i, chainLength)
		}
	}

	// Выводим общую статистику
	fmt.Println("\nОбщая статистика:")
	if minChain == math.MaxInt32 {
		fmt.Println("Минимальная длина непустой цепочки: 0")
	} else {
		fmt.Printf("Минимальная длина непустой цепочки: %d\n", minChain)
	}
	fmt.Printf("Максимальная длина цепочки: %d\n", maxChain)
	fmt.Printf("Пустых ячеек: %d из %d\n", emptyBuckets, ht.size)
	fmt.Printf("Общее количество элементов: %d\n", totalElements)
	fmt.Printf("Коэффициент заполнения: %.2f\n", float64(totalElements)/float64(ht.size))
}
