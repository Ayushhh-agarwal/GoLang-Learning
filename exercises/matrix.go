package exercises

import (
	"encoding/json"
	"fmt"
)

// 	Question :
//	Create a struct called 'Matrix'. The Matrix struct has the following information:
//		number of rows of matrix
//		number of columns of matrix
//		elements of matrix (You can use 2D vector)
//
//	The Matrix struct has methods for each of the following:
//		get the number of rows
//		get the number of columns
//		set the elements of the matrix at a given position (i,j)
//		adding two matrices. (2nd matrix can be taken as input to the method)
//		print matrix structure as json
//	You can assume that the dimensions are correct for the multiplication and addition.
//
//													Expectation: Use structs, slices, methods

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]int
}

func NewMatrix(rows, cols int) *Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, cols)
	}
	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		Elements: elements,
	}
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetCols() int {
	return m.Cols
}

func (m *Matrix) SetElement(i, j, val int) {
	m.Elements[i][j] = val
}

func (m *Matrix) Add(other *Matrix) *Matrix {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		return nil
	}

	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

func (m *Matrix) PrintAsJSON() {
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func RunMatrix() {
	// making 1st matrix
	//{
	//	1,2,3
	//	4,5,6
	//}

	matrix1 := NewMatrix(2, 3)
	matrix1.SetElement(0, 0, 1)
	matrix1.SetElement(0, 1, 2)
	matrix1.SetElement(0, 2, 3)
	matrix1.SetElement(1, 0, 4)
	matrix1.SetElement(1, 1, 5)
	matrix1.SetElement(1, 2, 6)

	// making 2nd matrix
	//{
	//	7,8,9
	//	10,11,12
	//}
	matrix2 := NewMatrix(2, 3)
	matrix2.SetElement(0, 0, 7)
	matrix2.SetElement(0, 1, 8)
	matrix2.SetElement(0, 2, 9)
	matrix2.SetElement(1, 0, 10)
	matrix2.SetElement(1, 1, 11)
	matrix2.SetElement(1, 2, 12)

	fmt.Println("Matrix 1:")
	matrix1.PrintAsJSON()

	fmt.Println("Matrix 2:")
	matrix2.PrintAsJSON()

	sum := matrix1.Add(matrix2)
	if sum != nil {
		fmt.Println("Sum of matrices:")
		sum.PrintAsJSON()
	} else {
		fmt.Println("Matrices cannot be added.")
	}
}
