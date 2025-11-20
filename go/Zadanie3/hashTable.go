package main

import (
	"fmt"
	"strconv"
)

type ChainingHashTable struct {
	lists []*ForwardListTwo // Храним списки напрямую в срезе
	size  int
}

func NewChainingHashTable(tableSize int) *ChainingHashTable {
	lists := make([]*ForwardListTwo, tableSize)

	for i := 0; i < tableSize; i++ {
		lists[i] = newForwardListTwo() // Создаем список для каждой ячейки
	}

	return &ChainingHashTable{
		lists: lists,
		size:  tableSize,
	}
}

func (ht *ChainingHashTable) hash(key int) int {
	if key < 0 {
		key = -key
	}
	return key % ht.size
}

func (ht *ChainingHashTable) getList(index int) *ForwardListTwo {
	if index < 0 || index >= ht.size {
		return nil
	}
	return ht.lists[index] // Возвращаем существующий список
}

func (ht *ChainingHashTable) clearList(list *ForwardListTwo) {
	if list == nil {
		return
	}

	for list.head != nil {
		deleteNodeHeadTwo(list)
	}
}

func (ht *ChainingHashTable) Insert(key int) {
	index := ht.hash(key)
	list := ht.getList(index)
	if list != nil {
		addNodeTailTwo(list, strconv.Itoa(key))
	}
}

func (ht *ChainingHashTable) Search(key int) bool {
	index := ht.hash(key)
	list := ht.getList(index)
	if list == nil {
		return false
	}
	return findNodeIndexTwo(list, strconv.Itoa(key))
}

func (ht *ChainingHashTable) Remove(key int) {
	index := ht.hash(key)
	list := ht.getList(index)
	if list != nil {
		deleteNodeIndexTwo(list, strconv.Itoa(key))
	}
}

func (ht *ChainingHashTable) GetStatistics() (int, int, float64) {
	minLength := 1 << 30
	maxLength := 0
	totalElements := 0
	nonEmptyBuckets := 0

	for i := 0; i < ht.size; i++ {
		list := ht.getList(i)
		if list != nil {
			listSize := countNodesTwo(list)
			if listSize > 0 {
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

	if minLength == 1<<30 {
		minLength = 0
	}

	avgLength := 0.0
	if nonEmptyBuckets > 0 {
		avgLength = float64(totalElements) / float64(nonEmptyBuckets)
	}

	return minLength, maxLength, avgLength
}

func (ht *ChainingHashTable) PrintDetailedStatistics() {
	totalElements := 0
	emptyBuckets := 0
	maxChain := 0
	minChain := 1 << 30

	fmt.Println("Распределение цепочек:")
	for i := 0; i < ht.size; i++ {
		list := ht.getList(i)
		chainLength := 0
		if list != nil {
			chainLength = countNodesTwo(list)
		}

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

		if i < 10 {
			fmt.Printf("Ячейка %d: %d элементов\n", i, chainLength)
		}
	}

	if minChain == 1<<30 {
		minChain = 0
	}

	fmt.Println("\nОбщая статистика:")
	fmt.Printf("Минимальная длина непустой цепочки: %d\n", minChain)
	fmt.Printf("Максимальная длина цепочки: %d\n", maxChain)
	fmt.Printf("Пустых ячеек: %d из %d\n", emptyBuckets, ht.size)
	fmt.Printf("Общее количество элементов: %d\n", totalElements)
	fmt.Printf("Коэффициент заполнения: %.2f\n", float64(totalElements)/float64(ht.size))
}

// Добавляем метод GetList для использования в Set
func (ht *ChainingHashTable) GetList(index int) *ForwardListTwo {
	return ht.getList(index)
}
