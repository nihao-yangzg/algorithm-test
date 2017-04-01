package main

import (
    "fmt"
)

func find_max_subarray(array []int) int {
    var custom = array[0]
    var max = array[0]
    for i := 1; i < len(array); i++ {
         if custom < 0 {
             custom = 0
         }
        custom += array[i]
        if (max < custom) {
            max = custom
        }
    }
    return max
}

func find_max_subarray1(array []int) int{
    var max = array[0]
    var custom = array[0]
    for i := 0; i < len(array); i++ {
        custom = 0
        for j := i; j < len(array); j++ {
            custom += array[j]
            if custom > max {
               max = custom
            }
        }
    }
    return max


}
func main(){
    var array = []int {-2, 5, 3, -6, 4, -8, -16}
    //fmt.Println(find_max_subarray(array))
    fmt.Println(find_max_subarray1(array))
    
}
