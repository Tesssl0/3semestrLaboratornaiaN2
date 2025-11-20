// hash_table.go
package main

import (
	"fmt"
	"math"
	"strconv"
)

type ChainingHashTable struct {
	table []*ForwardListTwo
	size  int
}

func NewChainingHashTable(tableSize int) *ChainingHashTable {
	ht := &ChainingHashTable{
		table: make([]*ForwardListTwo, tableSize),
		size:  tableSize,
	}

	// Инициализация пустых списков для каждой ячейки
	for i := 0; i < tableSize; i++ {
		ht.table[i] = newForwardListTwo()
	}

	return ht
}

func (ht *ChainingHashTable) hash(key int) int {
	return int(math.Abs(float64(key))) % ht.size
}

func (ht *ChainingHashTable) getList(index int) *ForwardListTwo {
	if index < 0 || index >= ht.size {
		return nil
	}
	return ht.table[index]
}

func (ht *ChainingHashTable) clearList(list *ForwardListTwo) {
	if list == nil {
		return
	}

	// Удаляем все узлы списка
	for list.head != nil {
		list.deleteNodeHead()
	}
}

func (ht *ChainingHashTable) Insert(key int) {
	index := ht.hash(key)
	list := ht.getList(index)
	if list != nil {
		list.addNodeTail(strconv.Itoa(key))
	}
}

func (ht *ChainingHashTable) Search(key int) bool {
	index := ht.hash(key)
	list := ht.getList(index)
	if list == nil {
		return false
	}

	return list.findNodeIndex(strconv.Itoa(key))
}

func (ht *ChainingHashTable) Remove(key int) {
	index := ht.hash(key)
	list := ht.getList(index)
	if list != nil {
		list.deleteNode(strconv.Itoa(key))
	}
}

func (ht *ChainingHashTable) GetAllElements() []int {
	var elements []int
	for i := 0; i < ht.size; i++ {
		list := ht.table[i]
		if list != nil {
			current := list.head
			for current != nil {
				value, err := strconv.Atoi(current.node)
				if err == nil {
					elements = append(elements, value)
				}
				current = current.next
			}
		}
	}
	return elements
}

func (ht *ChainingHashTable) GetStatistics() (int, int, float64) {
	minLength := math.MaxInt32
	maxLength := 0
	totalElements := 0
	nonEmptyBuckets := 0

	for i := 0; i < ht.size; i++ {
		list := ht.getList(i)
		if list != nil {
			listSize := list.countNodes()
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

	if minLength == math.MaxInt32 {
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
	minChain := math.MaxInt32

	fmt.Println("Распределение цепочек:")

	for i := 0; i < ht.size; i++ {
		list := ht.getList(i)
		chainLength := 0
		if list != nil {
			chainLength = list.countNodes()
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
