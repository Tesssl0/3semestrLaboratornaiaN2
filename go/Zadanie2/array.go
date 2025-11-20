// array.go
package main

import (
	"fmt"
	"strings"
)

type Array struct {
	data     []string
	size     int
	capacity int
}

func initArray(initialCapacity int) *Array {
	return &Array{
		data:     make([]string, initialCapacity),
		size:     0,
		capacity: initialCapacity,
	}
}

func (arr *Array) destroy() {
	arr.data = nil
	arr.size = 0
	arr.capacity = 0
}

func (arr *Array) resize() {
	newCapacity := arr.capacity * 2
	if newCapacity == 0 {
		newCapacity = 1
	}
	newData := make([]string, newCapacity)
	copy(newData, arr.data)
	arr.data = newData
	arr.capacity = newCapacity
}

func (arr *Array) add(value string) {
	if arr.size == arr.capacity {
		arr.resize()
	}
	arr.data[arr.size] = value
	arr.size++
}

func (arr *Array) addAt(index int, value string) {
	if arr.size == arr.capacity {
		arr.resize()
	}
	for i := arr.size; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	arr.data[index] = value
	arr.size++
}

func (arr *Array) get(index int) string {
	if index < 0 || index >= arr.size {
		return ""
	}
	return arr.data[index]
}

func (arr *Array) remove(index int) {
	if index < 0 || index >= arr.size {
		return
	}
	for i := index; i < arr.size-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.size--
}

func (arr *Array) set(index int, value string) {
	if index < 0 || index >= arr.size {
		return
	}
	arr.data[index] = value
}

func (arr *Array) length() int {
	return arr.size
}

func (arr *Array) print() {
	fmt.Println(strings.Join(arr.data[:arr.size], " "))
}

func (arr *Array) find(value string) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == value {
			return i
		}
	}
	return -1
}
