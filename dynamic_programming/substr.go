package main

import (
    "fmt"
)
/**
 *求解子串位置
 */
func substr(t string, s string) int {
    var t_l = len(t)
    var s_l = len(s)
    var result = -1
    for i := 0; i < t_l - s_l; i++ {
        j := i
        k := 0
        for j < t_l && k < s_l {

            if t[j] == s[k] {
                j++
                k++
            } else {
                break
            }
        }
        if k == s_l {
            result = i
            break
        } else {
            result = -1
        }
    }
    return result
}

func main(){
    var t = "abcdefdg"
    var s = "cdefs"
    fmt.Println(substr(t,s))

}
