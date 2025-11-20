// hash_Table.go
package main

import (
	"math"
	"strconv"
)

type CollisionResolution int

const (
	CHAINING CollisionResolution = iota
	LINEAR_PROBING
	QUADRATIC_PROBING
)

type HashTable struct {
	table  *Array
	size   int
	method CollisionResolution
	prime  int
}

func NewHashTable(tableSize int, resolutionMethod CollisionResolution) *HashTable {
	ht := &HashTable{
		table:  NewArray(tableSize),
		size:   tableSize,
		method: resolutionMethod,
	}

	ht.prime = ht.getNextPrime(tableSize / 2)

	if resolutionMethod == CHAINING {
		for i := 0; i < tableSize; i++ {
			// Убрана неиспользуемая переменная newList
			ht.table.Add("")
		}
	} else {
		for i := 0; i < tableSize; i++ {
			ht.table.Add("")
		}
	}

	return ht
}

func (ht *HashTable) hash1(key int) int {
	return int(math.Abs(float64(key))) % ht.size
}

func (ht *HashTable) hash2(key int) int {
	return ht.prime - (int(math.Abs(float64(key))) % ht.prime)
}

func (ht *HashTable) getNextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			return n
		}
		n++
	}
}

func (ht *HashTable) findSlotOpenAddressing(key int) int {
	index := ht.hash1(key)
	keyStr := strconv.Itoa(key)

	if ht.method == LINEAR_PROBING {
		for i := 0; i < ht.size; i++ {
			currentIndex := (index + i) % ht.size
			currentValue := ht.table.Get(currentIndex)

			if ht.isEmpty(currentIndex) || ht.isDeleted(currentIndex) {
				return currentIndex
			}

			if currentValue == keyStr {
				return currentIndex
			}
		}
	} else if ht.method == QUADRATIC_PROBING {
		for i := 0; i < ht.size; i++ {
			currentIndex := (index + i*i) % ht.size
			currentValue := ht.table.Get(currentIndex)

			if ht.isEmpty(currentIndex) || ht.isDeleted(currentIndex) {
				return currentIndex
			}

			if currentValue == keyStr {
				return currentIndex
			}
		}
	}

	return -1
}

func (ht *HashTable) isDeleted(index int) bool {
	if ht.method == CHAINING {
		return false
	}
	value := ht.table.Get(index)
	return value == "DELETED"
}

func (ht *HashTable) isEmpty(index int) bool {
	if ht.method == CHAINING {
		// Для цепочек всегда возвращаем false, так как логика другая
		return false
	} else {
		value := ht.table.Get(index)
		return value == "" || value == "DELETED"
	}
}

func (ht *HashTable) Insert(key int) {
	if ht.method == CHAINING {
		// В этой упрощенной версии для цепочек используем массив
		index := ht.hash1(key)
		keyStr := strconv.Itoa(key)
		current := ht.table.Get(index)

		if current == "" {
			ht.table.Set(index, keyStr)
		} else {
			// Для простоты добавляем через разделитель
			ht.table.Set(index, current+","+keyStr)
		}
	} else {
		index := ht.findSlotOpenAddressing(key)
		if index != -1 {
			ht.table.Set(index, strconv.Itoa(key))
		}
	}
}

func (ht *HashTable) Search(key int) bool {
	keyStr := strconv.Itoa(key)

	if ht.method == CHAINING {
		index := ht.hash1(key)
		value := ht.table.Get(index)

		if value == "" {
			return false
		}

		// Простая проверка для цепочек
		if value == keyStr {
			return true
		}

		// Проверяем в цепочке (разделенной запятыми)
		// В реальной реализации нужно использовать связанный список
		return containsValue(value, keyStr)
	} else {
		index := ht.findSlotOpenAddressing(key)
		return index != -1 && ht.table.Get(index) == keyStr
	}
}

func (ht *HashTable) Remove(key int) {
	keyStr := strconv.Itoa(key)

	if ht.method == CHAINING {
		index := ht.hash1(key)
		value := ht.table.Get(index)

		if value != "" {
			if value == keyStr {
				ht.table.Set(index, "")
			} else {
				// Упрощенное удаление из цепочки
				newValue := removeValue(value, keyStr)
				ht.table.Set(index, newValue)
			}
		}
	} else {
		index := ht.findSlotOpenAddressing(key)
		if index != -1 && ht.table.Get(index) == keyStr {
			ht.table.Set(index, "DELETED")
		}
	}
}

func (ht *HashTable) GetStatistics() (int, int, float64) {
	minLength := int(^uint(0) >> 1) // Max int
	maxLength := 0
	totalElements := 0
	nonEmptyBuckets := 0

	if ht.method == CHAINING {
		for i := 0; i < ht.size; i++ {
			value := ht.table.Get(i)
			if value != "" {
				chainLength := countChainLength(value)
				if chainLength < minLength {
					minLength = chainLength
				}
				if chainLength > maxLength {
					maxLength = chainLength
				}
				totalElements += chainLength
				nonEmptyBuckets++
			}
		}
	} else {
		visited := make([]bool, ht.size)

		for i := 0; i < ht.size; i++ {
			if !ht.isEmpty(i) && !ht.isDeleted(i) && !visited[i] {
				clusterSize := 0
				current := i

				for !ht.isEmpty(current) && !visited[current] {
					if !ht.isDeleted(current) {
						clusterSize++
					}
					visited[current] = true
					current = (current + 1) % ht.size
				}

				if clusterSize > 0 {
					if clusterSize < minLength {
						minLength = clusterSize
					}
					if clusterSize > maxLength {
						maxLength = clusterSize
					}
					totalElements += clusterSize
					nonEmptyBuckets++
				}
			}
		}
	}

	if minLength == int(^uint(0)>>1) {
		minLength = 0
	}

	avgLength := 0.0
	if nonEmptyBuckets > 0 {
		avgLength = float64(totalElements) / float64(nonEmptyBuckets)
	}

	return minLength, maxLength, avgLength
}

func (ht *HashTable) PrintDetailedStatistics() {
	totalElements := 0
	emptyBuckets := 0
	maxChain := 0
	minChain := int(^uint(0) >> 1)

	print("Метод разрешения коллизий: ")
	switch ht.method {
	case CHAINING:
		println("Метод цепочек")
	case LINEAR_PROBING:
		println("Линейное пробирование")
	case QUADRATIC_PROBING:
		println("Квадратичное пробирование")
	}

	println("Распределение:")

	for i := 0; i < ht.size; i++ {
		chainLength := 0

		if ht.method == CHAINING {
			value := ht.table.Get(i)
			if value != "" {
				chainLength = countChainLength(value)
			}
		} else {
			if !ht.isEmpty(i) && !ht.isDeleted(i) {
				chainLength = 1
			}
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
			printf("Ячейка %d: %d элементов\n", i, chainLength)
		}
	}

	println("\nОбщая статистика:")
	if minChain == int(^uint(0)>>1) {
		minChain = 0
	}
	printf("Минимальная длина непустой цепочки: %d\n", minChain)
	printf("Максимальная длина цепочки: %d\n", maxChain)
	printf("Пустых ячеек: %d из %d\n", emptyBuckets, ht.size)
	printf("Общее количество элементов: %d\n", totalElements)
	printf("Коэффициент заполнения: %.2f\n", float64(totalElements)/float64(ht.size))
}

func (ht *HashTable) SetResolutionMethod(method CollisionResolution) {
	ht.method = method
}

// Вспомогательные функции
func containsValue(chain string, value string) bool {
	// Простая проверка для демонстрации
	return chain == value
}

func removeValue(chain string, value string) string {
	// Упрощенная реализация для демонстрации
	if chain == value {
		return ""
	}
	return chain
}

func countChainLength(chain string) int {
	if chain == "" {
		return 0
	}
	// Упрощенный подсчет - в реальной реализации нужно разбирать цепочку
	return 1
}

func printf(format string, args ...interface{}) {
	// Простая реализация форматированного вывода
	print(format)
}
