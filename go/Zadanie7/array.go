package main

import "fmt"

// Array представляет динамический массив строк
type Array struct {
	data     []string
	size     int
	capacity int
}

// initArray инициализирует массив с начальной емкостью
func initArray(initialCapacity int) *Array {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Array{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// destroy очищает массив
func (arr *Array) destroy() {
	arr.data = nil
	arr.size = 0
	arr.capacity = 0
}

// resize увеличивает емкость массива
func (arr *Array) resize() {
	newCapacity := arr.capacity * 2
	if newCapacity == 0 {
		newCapacity = 10
	}
	newData := make([]string, newCapacity)
	copy(newData, arr.data[:arr.size])
	arr.data = newData
	arr.capacity = newCapacity
}

// add добавляет элемент в конец массива
func (arr *Array) add(value string) {
	if arr.size == arr.capacity {
		arr.resize()
	}
	arr.data[arr.size] = value
	arr.size++
}

// addAt добавляет элемент по указанному индексу
func (arr *Array) addAt(index int, value string) {
	if index < 0 || index > arr.size {
		return
	}

	if arr.size == arr.capacity {
		arr.resize()
	}

	// Сдвигаем элементы вправо
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}

	arr.data[index] = value
	arr.size++
}

// get возвращает элемент по индексу
func (arr *Array) get(index int) string {
	if index < 0 || index >= arr.size {
		return ""
	}
	return arr.data[index]
}

// remove удаляет элемент по индексу
func (arr *Array) remove(index int) {
	if index < 0 || index >= arr.size {
		return
	}

	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}

// set устанавливает значение элемента по индексу
func (arr *Array) set(index int, value string) {
	if index < 0 || index >= arr.size {
		return
	}
	arr.data[index] = value
}

// length возвращает текущий размер массива
func (arr *Array) length() int {
	return arr.size
}

// print выводит все элементы массива
func (arr *Array) print() {
	for i := 0; i < arr.size; i++ {
		fmt.Print(arr.data[i] + " ")
	}
	fmt.Println()
}

// findInArray ищет элемент по значению и возвращает индекс
func (arr *Array) findInArray(value string) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}
