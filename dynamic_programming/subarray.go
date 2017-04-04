package main

import "fmt"

func find_max_subarray(array []int) (int, int){
    var result [][2]int
    for i := 0; i < len(array); i++ {
        var temp [2]int
        for j := 0; j < 2; j++ {
            temp[j] = 0
        }
        result = append(result, temp)
    }
    result[0][0] = 0
    result[0][1] = 0
    
    var start, end int
    var max int
    for i := 1; i < len(array); i++ {
        if array[i] > array[i - 1] {
            end++
            if (end - start > max) {
                result[i][0] = start
                result[i][1] = end 
                max = end - start
            } else {
                result[i][0] = result[i-1][0]
                result[i][1] = result[i-1][1]
            }
        } else {
            start = i
            end = i
            result[i][0] = result[i-1][0]
            result[i][1] = result[i-1][1]
        }
    }
    fmt.Println(max)
    fmt.Println(result)
    return result[len(array) - 1][0], result[len(array) - 1][1]
}

func main() {
    var array = []int {1,1,4,3,6,5,7}
    fmt.Println(find_max_subarray(array))
}