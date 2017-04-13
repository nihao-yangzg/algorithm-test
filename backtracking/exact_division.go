package main

import (
    "fmt"
    "strings"
)

/**
 *查找能被整除的数量: 给定一个数字m（不大于18位），但是数字中的某些位不可见，如88888X88， 'X'为不可见位（可为多个），令给定一个整数n,求有多少种X的可能值使得m能够被整除。
 * 当不可见位数交大时,计算速度很慢.
 */

var result = 0
var origin = "999XX9XX999X"
func findExactDivision(dataStr string, size int, division int) {
    if size == 0 {
        length := len(dataStr)
        var data int64
        for i := 0; i < length; i++ {
            data = data * 10 + int64(dataStr[i] - '0')
        }
        if data % int64(division) == 0 {
            result ++
        }
        return
    }
    var begin int
    
    for h := 0; h < size; h++ {
        if dataStr[0] == 'X' {
            begin = 1
        }
        index := strings.Index(dataStr, "X")
        for i := begin; i < 10; i ++ {
            str := []byte(dataStr)
            str[index] = byte(i + '0')             
            findExactDivision(string(str), size-1, division)
            // fmt.Println(string(str))
        }
    }
    
}

func main(){
    var division = 3
    var size int
    for i := 0; i < len(origin); i++{
        if origin[i] == 'X' {
            size ++ 
        }
    }
    findExactDivision(origin, size, division)
    fmt.Println(result)
}