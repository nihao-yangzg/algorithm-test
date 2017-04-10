package main
/**
 * 找出所有找零的方式
 */

import (
    "fmt"
)

var coins = []int{2,3,5,7}
var target = 16
var result [][]int
func find(subrst []int, summary int, index int){
    for i := index; i < len(coins); i++ {
        subrst = append(subrst, coins[i])
        sum := summary + coins[i]
        if sum == target {
            result = append(result, subrst)
        } else if sum > target {
            subrst = subrst[0 : len(subrst) - 1]
            continue
        } else {
            find(subrst, sum, i)
        }
        temp := []int{}
        for j := 0; j < len(subrst) - 1; j++ {
            temp = append(temp, subrst[j])
        }
        subrst = temp
        
    }
}

func main(){
    var subrst = []int{}
    find(subrst, 0, 0)
    for i := 0; i < len(result); i++ {
        fmt.Println(result[i])
    }
}