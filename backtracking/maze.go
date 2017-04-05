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
        //判断是否超出边界
        if x-1 < 0 || x + 1 >= length || y - 1 < 0 || y + 1 >= length {
            continue
        }

        switch k {
            case 0: //上移
                if the_maze[x-1][y] != 0 {
                    start[0] = x-1
                    start[1] = y
                    ok = true
                }
                break
            case 1: //右移
                if the_maze[x][y+1] != 0 {
                    start[0] = x
                    start[1] = y+1
                    ok = true
                }
                break
            case 2: //下移
                if the_maze[x+1][y] != 0 {
                    start[0] = x+1
                    start[1] = y
                    ok = true
                    
                }
                break
            case 3: //左移
                if the_maze[x][y-1] != 0 {
                    start[0] = x
                    start[1] = y-1
                    ok = true
                }
                break
        }
       
        if ok {
            //访问过的元素置0，防止重复访问
            the_maze[x][y] = 0

            maze_path(start)
        }
        if found {
            //保存路径结果
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