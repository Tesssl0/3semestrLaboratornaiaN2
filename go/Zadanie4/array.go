package main

import "fmt"

// Array представляет динамический массив строк
type Array struct {
	data     []string
	size     int
	capacity int
}

// InitArray инициализирует массив
func InitArray(initialCapacity int) *Array {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Array{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

// Destroy очищает массив
func (arr *Array) Destroy() {
	arr.data = nil
	arr.size = 0
	arr.capacity = 0
}

// Resize увеличивает размер массива
func (arr *Array) Resize() {
	arr.capacity *= 2
	newData := make([]string, arr.capacity)
	copy(newData, arr.data)
	arr.data = newData
}

// Add добавляет элемент в конец
func (arr *Array) Add(value string) {
	if arr.size == arr.capacity {
		arr.Resize()
	}
	arr.data[arr.size] = value
	arr.size++
}

// AddAt добавляет элемент по индексу
func (arr *Array) AddAt(index int, value string) {
	if index < 0 || index > arr.size {
		return
	}
	if arr.size == arr.capacity {
		arr.Resize()
	}
	// Сдвигаем элементы вправо
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = value
	arr.size++
}

// Get возвращает элемент по индексу
func (arr *Array) Get(index int) string {
	if index < 0 || index >= arr.size {
		return ""
	}
	return arr.data[index]
}

// Remove удаляет элемент по индексу
func (arr *Array) Remove(index int) {
	if index < 0 || index >= arr.size {
		return
	}
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}

// Set устанавливает значение элемента по индексу
func (arr *Array) Set(index int, value string) {
	if index < 0 || index >= arr.size {
		return
	}
	arr.data[index] = value
}

// Length возвращает текущий размер массива
func (arr *Array) Length() int {
	return arr.size
}

// Print выводит все элементы массива
func (arr *Array) Print() {
	for i := 0; i < arr.size; i++ {
		fmt.Print(arr.data[i] + " ")
	}
	fmt.Println()
}

// FindInArray ищет элемент по значению и возвращает индекс
func (arr *Array) FindInArray(value string) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}
