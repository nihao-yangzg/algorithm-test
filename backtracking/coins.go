package main
/**
 * 找出所有找零的方式
 */

import (
    "fmt"
)

var coins = []int{2,3,5,7}
var target = 15
var result [][]int
func find(subrst []int, summary int, index int, max int){
    for i := index; i < len(coins); i++ {
        subrst = append(subrst, coins[i])
        sum := summary + coins[i]
        if sum == target {
            result = append(result, subrst)
        } else if sum > target {
            subrst = subrst[0 : len(subrst) - 1]
            continue
        } else {
            find(subrst, sum, i, max)
        }
        //开辟新的空间，并将原来的元素拷贝进入新空间，最后一个元素忽略
        temp := make([]int, len(subrst) - 1, max)
        for j := 0; j < len(subrst) - 1; j++ {
            // temp = append(temp, subrst[j])
            temp[j] = subrst[j]
        }
        subrst = temp
        
    }
}

func main(){
    var min = coins[0]
    for i := 1; i < len(coins); i++ {
        if coins[i] < min {
            min = coins[i]
        }
    }
    var subrstMaxLength = target / min
    var subrst = []int{} 
    find(subrst, 0, 0, subrstMaxLength)
    for i := 0; i < len(result); i++ {
        fmt.Println(result[i])
    }
}