package main

import (
	"fmt"
)

type Matrix struct {
	rows, cols int
	data       []float64
}

func NewMatrix(rows, cols int) *Matrix {
	if rows <= 0 || cols <= 0 {
		rows, cols = 1, 1
	}
	data := make([]float64, rows*cols)
	return &Matrix{rows, cols, data}
}

func (m *Matrix) Get(i, j int) float64 {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return 0
	}
	return m.data[i*m.cols+j]
}

func (m *Matrix) Set(i, j int, value float64) {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return
	}
	m.data[i*m.cols+j] = value
}

func (m *Matrix) Print() {
	for i := 0; i < m.rows; i++ {
		fmt.Print("| ")
		for j := 0; j < m.cols; j++ {
			value := m.Get(i, j)
			if value == 0 {
				fmt.Print("_")
			} else {
				fmt.Printf("%f", value)
			}
			fmt.Print(" ")
		}
		fmt.Println("|")
	}
}

func main() {
	new2DMatrix := NewMatrix(3, 3)
	new2DMatrix.Set(1, -1, 777)
	new2DMatrix.Set(0, 0, 1.1)
	new2DMatrix.Set(0, 1, 2878676.907079709)
	new2DMatrix.Set(0, 2, 999)
	new2DMatrix.Set(2, 2, -22)

	fmt.Printf("Get: %f\n", new2DMatrix.Get(2, 2))

	fmt.Printf("Matrix:\n")
	new2DMatrix.Print()

	fmt.Printf("\nCringe Matrix: ")
	new2DMatrix = NewMatrix(0, 10)
	new2DMatrix.Print()
}
