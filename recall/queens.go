package main

import (
    "fmt"
)
/**
 * 八皇后问题求解过程
 */
func search(result []int, cur int, num *int) {
    if cur == 8 {
       (*num)++
       fmt.Println(result)
       return
    }

    for i := 0; i < 8; i++ {
        result[cur] = i
        ok := true
        for j := 0; j < cur; j++ {
            if (result[cur] == result[j]) || (result[cur] - result[j] == cur - j) || (result[cur]-result[j] == j - cur) {
                ok = false
                break
            }
        }
        if ok {
           c := cur + 1
           search(result,c, num)
        }
    }
}

func main(){
    var result = []int{0,0,0,0,0,0,0,0}
    var num int
    num = 0
    search(result, 0, &num)
    fmt.Println(num)
}
