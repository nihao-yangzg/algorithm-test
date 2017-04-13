package main

import (
    "fmt"
)
/**
 * 将GBGB...GGBB调整顺序,调整后的顺序为GGG...BB或BB...GG,求调整次数最少为多少
 */
func adapter(queue string) int {
    var strlen = len(queue)
    var result = make([][2]int, strlen)
    var Gsize,Bsize int
    if queue[0] == 'G' {
        Gsize = 1
    } else {
        Bsize = 1
    }
    for i := 1; i < strlen; i++ {
        if queue[i] == 'G' {
        
            result[i][0] = result[i-1][0] + Bsize
            result[i][1] = result[i-1][1]
            Gsize++            
        } else {
            result[i][0] = result[i-1][0]
            result[i][1] = result[i-1][1] + Gsize
            Bsize++ 
        }
    }
    if result[strlen-1][0] > result[strlen-1][1] {
        return result[strlen-1][1]
    } 
    return result[strlen-1][0]
}
func main(){
    var queue = []string{"GBGBGG", "GGBBGGBGBG", "BBBGG", "GBGBGBGB"}
    for i := 0; i < len(queue); i++ {
        fmt.Println(adapter(queue[i]))
    }
}