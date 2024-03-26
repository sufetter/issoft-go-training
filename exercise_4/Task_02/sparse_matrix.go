package main

import (
	"fmt"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type SparseMatrix[T Number] struct {
	data map[int]map[int]T
}

func NewSparseMatrix[T Number]() *SparseMatrix[T] {
	return &SparseMatrix[T]{
		data: make(map[int]map[int]T),
	}
}

func (sm *SparseMatrix[T]) Set(row, col int, value T) {
	if row < 0 || col < 0 {
		panic("Invalid indices: row and col must be non-negative")
	}
	if _, ok := sm.data[row]; !ok {
		sm.data[row] = make(map[int]T)
	}
	sm.data[row][col] = value
}

func (sm *SparseMatrix[T]) Get(row, col int) T {
	if row < 0 || col < 0 {
		panic("Invalid indices: row and col must be non-negative")
	}
	if _, ok := sm.data[row]; !ok {
		var zero T
		return zero
	}
	return sm.data[row][col]
}

func main() {
	sm := NewSparseMatrix[float64]()

	sm.Set(1, 1, 5.0)
	sm.Set(2, 3, 8.0)

	fmt.Println(sm.Get(1, 1))
	fmt.Println(sm.Get(2, 3))
	fmt.Println(sm.Get(3, 3))

	// Try to set an element with an invalid type (string)
	// This will cause a compilation error
	// sm.Set(1, 2, "invalid")
}
