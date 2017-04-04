package main

import (
    "fmt"
)
/**
 *读塔问题
 */
func count(tower [][]int) int {
    for i := len(tower) - 1; i >= 0; i-- {
        for j := 0; j < i; j++ {
            if tower[i][j] > tower[i][j+1] {
                tower[i-1][j] +=tower[i][j]
            } else {
                tower[i-1][j] +=tower[i][j+1]
           }
        }
    }
    return tower[0][0]

}

func main() {
    var tower = [][]int{
                 {9,0,0,0},
                 {12,15,0,0},
                 {10,6,8,0},
                 {2,18,9,5},
                }
    fmt.Println(count(tower))

}
