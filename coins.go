package main
import "fmt"

func find_num(coins []int, size int, amount int) int{
    var result [][]int
    for i:=0; i<size + 1; i++ {
        var temp []int      
        for j:=0; j<amount + 1; j++ {
            temp = append(temp,0)
        }
        result = append(result,temp)
    }
    for i:=0; i<size + 1; i++ {
        result[i][0] = 1
    }


    for i:=1; i<size + 1; i++ {
        for j := 1; j < amount + 1; j++ {
            for k := 0; k < j/coins[i-1] + 1; k++ {
                result[i][j] += result[i-1][j - k*coins[i-1]]
            }
        }
    }
    fmt.Println(result)
    
    return result[size][amount]
}

func main() {
    coins:= []int{1,2,5}
    size:= len(coins)
    amount := 5
    fmt.Println(find_num(coins, size, amount))
}
