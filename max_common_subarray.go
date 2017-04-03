package main

import (
    "fmt"
)
/**
 * 求最大公共子串长度
 */

func max_common_subarray(array1 []int, array2 []int) int {
    var result [][]int
    var len1 = len(array1)
    var len2 = len(array2)
    for i := 0; i < len1 + 1; i++ {
        var temp []int
        for j := 0 ; j < len2 + 1; j++ {
            temp = append(temp, 0)
        }
        result = append(result, temp)
    }
    //动态规划过程

    for i := 1; i < len1 + 1; i++ {
        for j := 1; j < len2 + 1; j++ {
            if array1[i-1] == array2[j-1] {
                result[i][j] = result[i-1][j] + 1
            }
            if result[i][j] < result[i][j-1] {
                result[i][j] = result[i][j-1]
            }
            if result[i][j] < result[i-1][j] {
                result[i][j] = result[i-1][j]
            }
        }
    }
    return result[len1][len2]

}


func main(){

    //var array1 = []int {2,3,1,5,4}
    //var array2 = []int {4,2,1,3,5}
    var array1 = []int {2,3,5,4,1,6}
    var array2 = []int {4,6,5,1,2}
    fmt.Println(max_common_subarray(array1, array2))
}
