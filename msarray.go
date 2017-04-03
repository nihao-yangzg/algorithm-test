package main

import "fmt"

/**
 * 求最长递增子串
 */
func max_sub(array []int) int {
    var result []int
    var length = len(array)
    for i := 0; i < length; i++ {
        result = append(result, 0)
    }
    //动态规划过程
    for i := 0; i < length; i++ {
        result[i] = 1
        for j := 0; j < i; j++ {
            if (array[i] > array[j] && result[i] <= result[j]){
                 result[i] = result[j] + 1
            }
        }
    }
    var max = result[0]
    for i := 1; i < length; i++ {
        if max < result[i] {
            max = result[i]
        }
    }
    return max
}

func main() {
     var array = []int {1,2,3,2,4,5,8,3}
     fmt.Println(max_sub(array))
}

