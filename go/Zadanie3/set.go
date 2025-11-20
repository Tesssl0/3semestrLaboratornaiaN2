package main

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
