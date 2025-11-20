package main

import "fmt"

type Array struct {
	data     []string
	size     int
	capacity int
}

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

func destroyArray(arr *Array) {
	arr.data = nil
	arr.size = 0
	arr.capacity = 0
}

func resizeArray(arr *Array) {
	newCapacity := arr.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	newData := make([]string, newCapacity)
	copy(newData, arr.data[:arr.size])
	arr.data = newData
	arr.capacity = newCapacity
}

func addToArray(arr *Array, value string) {
	if arr.size == arr.capacity {
		resizeArray(arr)
	}
	arr.data[arr.size] = value
	arr.size++
}

func addAtArray(arr *Array, index int, value string) {
	if index < 0 || index > arr.size {
		return
	}
	if arr.size == arr.capacity {
		resizeArray(arr)
	}
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = value
	arr.size++
}

func getFromArray(arr *Array, index int) string {
	if index < 0 || index >= arr.size {
		return ""
	}
	return arr.data[index]
}

func removeFromArray(arr *Array, index int) {
	if index < 0 || index >= arr.size {
		return
	}
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}

func setInArray(arr *Array, index int, value string) {
	if index < 0 || index >= arr.size {
		return
	}
	arr.data[index] = value
}

func arrayLength(arr *Array) int {
	return arr.size
}

func printArray(arr *Array) {
	for i := 0; i < arr.size; i++ {
		fmt.Print(arr.data[i], " ")
	}
	fmt.Println()
}

func findInArray(arr *Array, value string) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}
