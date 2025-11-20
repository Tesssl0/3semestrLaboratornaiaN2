// stack.go
package main

import "fmt"

type Stack struct {
	arr      []string
	top      int
	capacity int
	silent   bool
}

func NewStack(capacity int, silent bool) *Stack {
	return &Stack{
		arr:      make([]string, capacity),
		capacity: capacity,
		top:      -1,
		silent:   silent,
	}
}

func (s *Stack) Find(value string) int {
	for i := 0; i <= s.top; i++ {
		if s.arr[i] == value {
			return i
		}
	}
	return -1
}

func (s *Stack) Print() {
	for i := 0; i <= s.top; i++ {
		fmt.Print(s.arr[i])
		if i < s.top {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func (s *Stack) Push(value string) {
	if s.top >= s.capacity-1 {
		if !s.silent {
			fmt.Println("Стек переполнен!")
		}
		return
	}
	s.top++
	s.arr[s.top] = value
	if !s.silent {
		fmt.Printf("Элемент %s добавлен в стек\n", value)
	}
}

func (s *Stack) Pop() {
	if s.top < 0 {
		if !s.silent {
			fmt.Println("Стек пуст!")
		}
		return
	}
	if !s.silent {
		fmt.Printf("Элемент %s удалён из стека\n", s.arr[s.top])
	}
	s.top--
}

func (s *Stack) Top() string {
	if s.top < 0 {
		if !s.silent {
			fmt.Println("Стек пуст!")
		}
		return ""
	}
	return s.arr[s.top]
}

func (s *Stack) IsEmpty() bool {
	return s.top < 0
}

func (s *Stack) TopIndex() int {
	return s.top
}

func (s *Stack) GetAt(index int) string {
	if index < 0 || index > s.top {
		return ""
	}
	return s.arr[index]
}

func (s *Stack) Size() int {
	return s.top + 1
}
