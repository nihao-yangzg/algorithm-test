package main

/**
 * 数独求解
 */

import (
    "fmt"
)

// var origin_sudoku = [][]int{
//     {8,0,0,0,0,0,0,0,0},
//     {0,0,3,6,0,0,0,0,0},
//     {0,7,0,0,9,0,2,0,0},
//     {0,5,0,0,0,7,0,0,0},
//     {0,0,0,0,4,5,7,0,0},
//     {0,0,0,1,0,0,0,3,0},
//     {0,0,1,0,0,0,0,6,8},
//     {0,0,8,5,0,0,0,1,0},
//     {0,9,0,0,0,0,4,0,0},
// }
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

var origin_sudoku = [][]int{
    {0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0},
    {0,0,0,0,2,0,0,0,0},
    {0,1,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,5,0,0},
    {0,0,0,0,0,0,0,0,0},
    {0,0,0,0,0,0,0,0,0},
    
}
var found = false

func solve_sudoku(sudoku [][]int, index int) {
    if index == 81 {
        found = true
        return
    }
    if sudoku[1][0] == 7 {
        for i :=0; i < 9; i++ {
            fmt.Println(sudoku[i])
        }
    }
    var ok = true
    var cur int
    row := index / 9
    col := index % 9
    max := row
    if row < col {
        max = col 
    }
    for i := 0; i < 9 - max; i++ {
        if sudoku[row][col] != 0 && origin_sudoku[row][col] != 0{
            // fmt.Println(index)
            // fmt.Println(sudoku[4])
            solve_sudoku(sudoku, index + 1)
            break
        }
        for k := 1; k < 10; k ++ {
            if (k <= cur) {
                continue
            }
            cur = k
            // if origin_sudoku[row][col] == 0 && cur <= sudoku[row][col] {
                // continue
            // }
            ok = true

            for m := 0; m < 9; m++ {
                
                if (m > col) && (origin_sudoku[row][m] == cur){
                    ok = false
                    break
                } else if (m < col) && (cur == sudoku[row][m]) {
                    ok = false
                    break
                } else if m == col {
                    // ok = false
                    continue
                }
            }
            if !ok {
                continue
            }
            for n := 0; n < 9; n++ {
                if (n < row) && (cur == sudoku[n][col]){
                    ok = false
                    break
                } else if (n > row) && (cur == origin_sudoku[n][col]) {
                    ok = false
                    break
                } else if row == n {
                    // ok = false
                    continue
                }
            }
            if !ok {
                continue
            }

            for p := 0; p < 9; p++ {
                irow := row / 3 * 3 + p / 3 
                icol := col / 3 * 3 + p % 3
                if (irow < row || icol < col) && cur == sudoku[irow][icol]{
                    ok = false
                    break
                } else if (irow > row || icol > col) && cur == origin_sudoku[irow][icol] {
                    ok = false
                    break
                } else if (irow == row) && (icol == col){
                    // ok = false
                    continue
                } 
            }
            if !ok {
                continue
            } else {
                fmt.Printf("(%d,%d) -> %d\n", row, col, cur)
                sudoku[row][col] = cur
                // fmt.Println(sudoku[1])
                // fmt.Println(index)
                
                // fmt.Println(sudoku[5])
                // fmt.Println(sudoku[row])
                solve_sudoku(sudoku, index + 1)
                // break
                if found {
                    return
                }
            }
        }
    }
}

func main(){
    var sudoku [][]int
    for i := 0; i < 9; i++ {
        var temp []int 
        for j :=0; j < 9; j++ {
            temp = append(temp, origin_sudoku[i][j])
        }
        sudoku = append(sudoku, temp)
        
    }
    solve_sudoku(sudoku, 0)
    for i :=0; i < 9; i++ {
        fmt.Println(sudoku[i])
    }
}