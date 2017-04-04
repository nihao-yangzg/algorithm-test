package main

import (
    "fmt"
)

/**
 * 无向连通图着色问题
 */

func color(cur int) int {
    var num = 0
    if cur == len(graph) {
        fmt.Println(result)
        return 1
    }
    for i := 0; i < len(colors); i++ {
        result[cur] = colors[i]
        ok := true
        for j := 0; j < len(graph); j++ {
            point := graph[cur][j]
            if point != 0 && result[cur] == result[j] {
                ok = false
                break
            }
        }
        if (ok) {
           next := cur + 1
           num += color(next)
        }
    }
    return num

}

var graph = [][]int{
        {0,1,1,1,0,0},
        {1,0,0,0,0,0},
        {1,0,0,1,1,1},
        {1,0,1,0,0,0},
        {0,0,1,0,0,0},
        {0,0,1,0,0,0},
    }

var colors = []int{1,2,3}
var result = []int{0,0,0,0,0,0}
func main(){
    num := color(0)
    fmt.Println(num)
}
