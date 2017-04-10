package main

/**
 * 数独求解
 */

import (
	"fmt"
)

var origin_sudoku = [][]int{
	{8, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 0, 0},
	{0, 7, 0, 0, 9, 0, 2, 0, 0},
	{0, 5, 0, 0, 0, 7, 0, 0, 0},
	{0, 0, 0, 0, 4, 5, 7, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 3, 0},
	{0, 0, 1, 0, 0, 0, 0, 6, 8},
	{0, 0, 8, 5, 0, 0, 0, 1, 0},
	{0, 9, 0, 0, 0, 0, 4, 0, 0},
}

// var origin_sudoku = [][]int{
//     {3,1,0,0,0,0,0,8,5},
//     {0,0,0,9,0,3,0,0,0},
//     {9,0,5,0,0,0,3,0,7},
//     {8,0,4,0,0,0,1,0,6},
//     {0,0,0,4,0,1,0,0,0},
//     {6,9,0,0,0,0,0,7,3},
//     {0,3,0,5,0,2,0,1,0},
//     {0,0,0,8,0,4,0,0,0},
//     {0,2,0,7,0,6,0,9,0},
//    }
var found = false

func solveSudoku(sudoku [][]int, index int) {
	if index == 81 {
		found = true
		return
	}

	var ok = true
	var cur, row, col, max int
	row = index / 9
	col = index % 9
	max = row
	if row < col {
		max = col
	}
	for i := 0; i < 9-max; i++ {
		if sudoku[row][col] != 0 && origin_sudoku[row][col] != 0 {
			solveSudoku(sudoku, index+1)
			break
		}
		for j := 1; j < 10; j++ {
			if j <= cur {
				continue
			}
			cur = j
			ok = true
			//对比与当前行数据是否重复
			for k := 0; k < 9; k++ {

				if (k > col) && (origin_sudoku[row][k] == cur) {
					ok = false
					break
				} else if (k < col) && (cur == sudoku[row][k]) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			//对比与当前列数据是否重复
			for k := 0; k < 9; k++ {
				if (k < row) && (cur == sudoku[k][col]) {
					ok = false
					break
				} else if (k > row) && (cur == origin_sudoku[k][col]) {
					ok = false
					break
				}
			}
			if !ok {
				continue
			}
			//对比当前九宫格内是否重复
			for k := 0; k < 9; k++ {
				irow := row/3*3 + k/3
				icol := col/3*3 + k%3
				iindex := irow*9 + icol
				if (iindex < index) && cur == sudoku[irow][icol] {
					ok = false
					break
				} else if (iindex > index) && cur == origin_sudoku[irow][icol] {
					ok = false
					break
				}
			}
			if !ok {
				continue
			} else {
				sudoku[row][col] = cur
				solveSudoku(sudoku, index+1)
				if found {
					return
				}
			}
		}
	}
}

func main() {
	var sudoku [][]int
	for i := 0; i < 9; i++ {
		var temp []int
		for j := 0; j < 9; j++ {
			temp = append(temp, origin_sudoku[i][j])
		}
		sudoku = append(sudoku, temp)

	}
	fmt.Printf("sudoku:\n")
	for i := 0; i < 9; i++ {
		fmt.Println(origin_sudoku[i])
	}
	solveSudoku(sudoku, 0)

	fmt.Printf("The Answer:\n")
	if found {
		for i := 0; i < 9; i++ {
			fmt.Println(sudoku[i])
		}
	} else {
		fmt.Printf("We can not find the Answer.\n")
	}
}