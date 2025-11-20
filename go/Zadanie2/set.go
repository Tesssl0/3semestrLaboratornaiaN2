// set.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Set struct {
	table     *ChainingHashTable
	tableSize int
}

func NewSet(size int) *Set {
	if size <= 0 {
		size = 100
	}
	return &Set{
		table:     NewChainingHashTable(size),
		tableSize: size,
	}
}

func (s *Set) Add(value int) bool {
	if s.table.Search(value) {
		return false
	}
	s.table.Insert(value)
	return true
}

func (s *Set) Remove(value int) bool {
	if !s.table.Search(value) {
		return false
	}
	s.table.Remove(value)
	return true
}

func (s *Set) Contains(value int) bool {
	return s.table.Search(value)
}

func (s *Set) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Ошибка: не удалось открыть файл %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			value, err := strconv.Atoi(line)
			if err != nil {
				fmt.Printf("Неверный формат данных в файле: %s\n", line)
				continue
			}
			s.table.Insert(value)
		}
	}

	return scanner.Err()
}

func (s *Set) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Ошибка: не удалось открыть файл для записи %s", filename)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// Получаем все элементы из хеш-таблицы
	elements := s.table.GetAllElements()

	for _, value := range elements {
		_, err := writer.WriteString(strconv.Itoa(value) + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
