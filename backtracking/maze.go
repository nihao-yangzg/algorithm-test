package main

/**
 * 迷宫问题
 */

import (
     "fmt"
)

func maze_path(start []int){
    var length = len(the_maze)
    var x = start[0]
    var y = start[1]
    if start[0] == out[0] && start[1] == out[1] {
        found = true
        return
    }
    for k := 0; k < 4; k++ {
        ok := false
        if x-1 < 0 || x + 1 >= length || y - 1 < 0 || y + 1 >= length {
            continue
        }


        switch k {
            case 0:
                if the_maze[x-1][y] != 0 {
                    start[0] = x-1
                    start[1] = y
                    ok = true
                }
                break
            case 1:
                if the_maze[x][y+1] != 0 {
                    start[0] = x
                    start[1] = y+1
                    ok = true
                }
                break
            case 2:
                if the_maze[x+1][y] != 0 {
                    start[0] = x+1
                    start[1] = y
                    ok = true
                    
                }
                break
            case 3:
                if the_maze[x][y-1] != 0 {
                    start[0] = x
                    start[1] = y-1
                    ok = true
                }
                break
        }
       
        if ok {
            pre[0] = x
            pre[1] = y
            // fmt.Println(start)
            // fmt.Println(the_maze[8][3])
            // fmt.Println(pre)
            the_maze[x][y] = 0
            maze_path(start)
        }
        if found {
            // i  := start[0]
            // j := start[1]
            result = append(result, []int{x,y})
            break
        }
    }
}

var the_maze = [][]int {
    {0,0,0,0,0,0,0,0,0,0,0},
    {0,1,1,1,1,1,0,1,1,1,0},
    {0,1,0,1,1,1,0,1,0,1,0},
    {1,1,0,1,1,1,1,1,0,1,0},
    {0,1,0,0,0,0,0,1,0,1,0},
    {0,1,0,1,0,1,1,1,0,1,0},
    {0,1,1,1,0,1,0,1,1,1,0},
    {0,0,0,0,0,1,0,1,1,1,0},
    {0,1,0,1,0,1,0,1,1,1,0},
    {0,1,1,1,1,1,0,1,1,1,0},
    {0,0,0,0,0,0,0,0,0,0,0},
}
var in = []int {8,3}
var out = []int {3,0}
var result [][]int
var found = false
var pre = []int{0,0}
func main(){
    maze_path(in)
    for i := len(result) - 1; i > 0; i-- {
        fmt.Print(result[i])
        fmt.Print("-->")
        fmt.Println(result[i-1])
    }
    fmt.Print(result[0])
    fmt.Print("-->")
    fmt.Println(out)

}