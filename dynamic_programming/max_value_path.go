package main

import (
    "fmt"
)


func max_path(values [][]int, end []int) int {
    var rvalue [][]int
    var pnum = len(values)
    var boundery = end[0]
    for i := 0; i < boundery + 1; i++ {
        var temp []int
        for j := 0; j < boundery + 1; j++ {
              temp = append(temp, 0)
        }
        rvalue = append(rvalue, temp)
    }
    for i := 0; i < pnum; i++ {
        t := uint8(values[i][0])
        p := uint8(values[i][1])
        rvalue[t][p] = values[i][2]
    }

    var result [][]int

    for i := 0; i < boundery + 1; i++ {
        var temp []int
        for j := 0; j < boundery + 1; j++ {
            temp = append(temp, 0)
        }
        result = append(result, temp)
    }

    for i := 1; i < boundery + 1; i++ {
        for j := 1; j < boundery + 1; j++ {
            if result[i-1][j] > result[i][j-1] {
               result[i][j] = result[i-1][j] + rvalue[i][j]
            } else {
               result[i][j] = result[i][j-1] + rvalue[i][j]
            }
        }
     }
     return result[boundery][boundery]
    


}

func main() {
    var values = [][]int {
        {1,3,3},{2,2,4},{2,6,12},{3,4,2},{4,5,7},{4,7,1},{5,1,2},{6,2,3},{6,4,1},{6,6,3},
    }
    var end = []int {8,8}
    fmt.Println(max_path(values, end))
         

}
