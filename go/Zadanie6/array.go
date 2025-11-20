// array.go
package main

type Array struct {
	data     []string
	size     int
	capacity int
}

func NewArray(initialCapacity int) *Array {
	if initialCapacity <= 0 {
		initialCapacity = 10
	}
	return &Array{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

func (arr *Array) Destroy() {
	arr.data = nil
	arr.size = 0
	arr.capacity = 0
}

func (arr *Array) resize() {
	newCapacity := arr.capacity * 2
	newData := make([]string, newCapacity)
	copy(newData, arr.data)
	arr.data = newData
	arr.capacity = newCapacity
}

func (arr *Array) Add(value string) {
	if arr.size == arr.capacity {
		arr.resize()
	}
	arr.data[arr.size] = value
	arr.size++
}

func (arr *Array) AddAt(index int, value string) {
	if arr.size == arr.capacity {
		arr.resize()
	}
	copy(arr.data[index+1:], arr.data[index:arr.size])
	arr.data[index] = value
	arr.size++
}

func (arr *Array) Get(index int) string {
	return arr.data[index]
}

func (arr *Array) Remove(index int) {
	copy(arr.data[index:], arr.data[index+1:arr.size])
	arr.size--
}

func (arr *Array) Set(index int, value string) {
	arr.data[index] = value
}

func (arr *Array) Length() int {
	return arr.size
}

func (arr *Array) Print() {
	for i := 0; i < arr.size; i++ {
		print(arr.data[i] + " ")
	}
	println()
}

func (arr *Array) Find(value string) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}
