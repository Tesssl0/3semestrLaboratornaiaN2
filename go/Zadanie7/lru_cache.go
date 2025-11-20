package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LRUCache представляет LRU кэш
type LRUCache struct {
	capacity  int
	usageList *ForwardListTwo
	cacheMap  *ChainingHashTable
}

// moveToFront перемещает элемент в начало списка
func (lru *LRUCache) moveToFront(key int) {
	keyStr := strconv.Itoa(key)
	current := lru.usageList.head

	// Ищем узел с данным ключом
	for current != nil {
		if strings.HasPrefix(current.node, keyStr+":") {
			nodeData := current.node
			deleteNodeIndexTwo(lru.usageList, nodeData)
			addNodeHeadTwo(lru.usageList, nodeData)
			break
		}
		current = current.next
	}
}

// removeLRU удаляет наименее используемый элемент
func (lru *LRUCache) removeLRU() {
	if lru.usageList.tail != nil {
		lruData := lru.usageList.tail.node
		colonPos := strings.Index(lruData, ":")
		if colonPos != -1 {
			key, _ := strconv.Atoi(lruData[:colonPos])
			lru.cacheMap.remove(key)
		}
		deleteNodeTailTwo(lru.usageList)
	}
}

// NewLRUCache создает новый LRU кэш
func NewLRUCache(cap int) *LRUCache {
	return &LRUCache{
		capacity:  cap,
		usageList: &ForwardListTwo{head: nil, tail: nil},
		cacheMap:  NewChainingHashTable(100),
	}
}

// set устанавливает значение в кэше
func (lru *LRUCache) set(key int, value int) {
	keyStr := strconv.Itoa(key)
	itemStr := keyStr + ":" + strconv.Itoa(value)

	// Если ключ уже существует в кэше
	if lru.cacheMap.search(key) {
		lru.moveToFront(key)
		// Обновляем значение в списке
		current := lru.usageList.head
		for current != nil {
			if strings.HasPrefix(current.node, keyStr+":") {
				current.node = itemStr
				break
			}
			current = current.next
		}
	} else {
		// Если кэш заполнен, удаляем LRU элемент
		if countNodesTwo(lru.usageList) >= lru.capacity {
			lru.removeLRU()
		}
		lru.cacheMap.insert(key)
		addNodeHeadTwo(lru.usageList, itemStr)
	}
}

// get получает значение из кэша
func (lru *LRUCache) get(key int) int {
	if lru.cacheMap.search(key) {
		lru.moveToFront(key)
		keyStr := strconv.Itoa(key)
		current := lru.usageList.head

		// Ищем элемент в списке и извлекаем значение
		for current != nil {
			if strings.HasPrefix(current.node, keyStr+":") {
				colonPos := strings.Index(current.node, ":")
				if colonPos != -1 {
					value, _ := strconv.Atoi(current.node[colonPos+1:])
					return value
				}
			}
			current = current.next
		}
	}
	return -1
}

// printCache выводит текущее состояние кэша
func (lru *LRUCache) printCache() {
	fmt.Print("LRU Cache: ")
	current := lru.usageList.head
	// Выводим все элементы от недавних к старым
	for current != nil {
		fmt.Print("[", current.node, "] ")
		current = current.next
	}
	fmt.Println()
}

// processLRUQueries обрабатывает пользовательские запросы к LRU кэшу
func processLRUQueries() {
	fmt.Print("Введите емкость кэша: ")
	var cap, Q int
	fmt.Scan(&cap)
	fmt.Print("Введите количество запросов: ")
	fmt.Scan(&Q)

	// Очищаем буфер ввода
	reader := bufio.NewReader(os.Stdin)

	cache := NewLRUCache(cap)
	fmt.Println("Введите запросы (SET x y или GET x):")

	// Счетчик для результатов GET запросов
	var results []string

	// Обрабатываем все запросы
	for i := 0; i < Q; i++ {
		query, _ := reader.ReadString('\n')
		query = strings.TrimSpace(query)

		parts := strings.Fields(query)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToUpper(parts[0])

		// Обработка команды SET
		if command == "SET" && len(parts) >= 3 {
			key, err1 := strconv.Atoi(parts[1])
			value, err2 := strconv.Atoi(parts[2])
			if err1 == nil && err2 == nil {
				cache.set(key, value)
				fmt.Printf("SET %d %d выполнено\n", key, value)
			}
		} else if command == "GET" && len(parts) >= 2 {
			// Обработка команды GET
			key, err := strconv.Atoi(parts[1])
			if err == nil {
				result := cache.get(key)
				results = append(results, strconv.Itoa(result))
				fmt.Printf("GET %d = %d\n", key, result)
			}
		}
		cache.printCache()
	}

	// Выводим все результаты GET запросов
	fmt.Print("Результаты: ")
	for i, result := range results {
		fmt.Print(result)
		if i < len(results)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
