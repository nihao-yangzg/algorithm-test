package main

import (
    "fmt"
)

/**
 * 寻找一个集合的全部子集合
 */

var array = []int{1,2,4,5,6}
var result [][]int
func subset(sub []int, index int) {
    for i := index; i < len(array); i++ {
        sub = append(sub, array[i])
        result = append(result, sub)
        subset(sub, i + 1)
        var temp []int
        for j := 0; j < len(sub) - 1 ; j++ {
            temp = append(temp, sub[j])
        }
        sub = temp
    } 
}

func main(){
    var temp []int
    subset(temp, 0)
    for i := 0; i < len(result); i++ {
        fmt.Println(result[i])
    }
}