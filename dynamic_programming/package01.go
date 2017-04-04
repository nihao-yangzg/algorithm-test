package main

/**
 *0-1背包问题求解过程
 */
import (
   "fmt"
)

func get_max(value []int, weigh []int, max_weigh int) int {
    var result [][]int
    var size = len(value)

    //初始化数组
    for i := 0; i <= max_weigh; i++ {
        var temp []int
        for j := 0; j < size; j++ {
            temp = append(temp, 0)
        }
        result = append(result, temp)
    }

    //动态规划过程
    //result[i][j]--> 限重为i，仅有前j个物品情况下的最大装包值。
    for i := 1; i <= max_weigh; i++ {
        for j := 1; j < size; j++ {
            if weigh[j] > i{
               result[i][j] = result[i][j-1]
            } else {
               result[i][j] = result[i-weigh[j]][j-1] + value[j]
               if result[i][j] < result[i][j-1]{
                   result[i][j] = result[i][j-1]
               }
            }
        }

    }
    //fmt.Println(result)
    return result[max_weigh][size-1]
}

func main() {
    var value = []int{2,2,4,3,5,10}
    var weigh = []int{3,2,2,1,3,6}
    var max_weigh = 6
    fmt.Println(get_max(value, weigh, max_weigh))
}
