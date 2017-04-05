package main

/**
 * 24 点问题
 */
import (
    "fmt"
)

func comm24(result int, index int) {
    if result == target && index == 4{
        found = true
       
        return
    }
    if index == 4 {
        return
    }
    var ok = true
    var cur, sindex, sresult int
    var symbol byte
    for i := 0; i < len(datum); i++ {
        ok = true
        for j := 0; j < index + 1; j++ {
            if selected[j] == datum[i] {
                ok = false
                break
            }
        }
        if ok {
            cur = datum[i]
            selected[index] = cur
            sindex = index + 1
            if (index == 0) {
                comm24(cur, sindex)
            } else {
                symbol = 0
                for k := 0; k < len(symbols); k++ {
                    symbol = symbols[k]
                    sym_result[index-1] = symbol
                    sresult = 0
                    
                    switch symbol {
                        case '+':
                            sresult = result + cur
                            break
                        case '-':
                            sresult = result - cur
                            break
                        case '*':
                            sresult = result * cur
                            break
                        case '/':
                            sresult = result / cur
                            break
                    }
                    
                    comm24(sresult, sindex)
                    if found {
                        break
                    }
                }
            }
        }
        if found {
            break;
        }
    }

    
}
var target = 20
var sym_result = [3]byte{}
var found = false
var datum = []int{1,2,3,7}
var symbols = "+-*/"
var selected = [4]int{-1,-1,-1,-1}

func output(i int, j int) string{
    if i == 1 || j == 0 {
        return fmt.Sprintf("(%d %c %d)", selected[i-1], sym_result[j], selected[i]) 
    }
    return fmt.Sprintf("(%s %c %s)", output(i-1,j-1), sym_result[j], fmt.Sprintf("%d",selected[i]))
}
func main() {
    comm24(0,0)
    if found {
        fmt.Printf("%s --> %d\n", output(3,2), target)
    } else {
        fmt.Println("not found")
    }
}