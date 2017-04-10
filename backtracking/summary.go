package main

import (
    "fmt"
)
/**
 * 找出给定数组中的子数组，该子数组满足其和等于要求的值
 */
var elements = []int{10,1,2,7,6,1,5}
var target = 8
func find(subset []int, index int, sum int) {
    for i := index; i < len(elements); i++ {
        subset = append(subset, elements[i])
        tsum := sum + elements[i]
        if tsum == target {
            result = append(result, subset)
        } else if tsum > target {
            subset = subset[0:len(subset)-1]
            continue
        } else {
            find(subset, i+1, tsum)
        }
        temp := []int{}
        for k := 0; k < len(subset) - 1; k++ {
            temp = append(temp, subset[k])
        }
        subset = temp
        
    }
}
var result [][]int

func main(){
    var subset = []int{}
    find(subset, 0, 0)
    for i := 0; i < len(result); i++ {
        fmt.Println(result[i])
    }    
}