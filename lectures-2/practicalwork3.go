package main 

import "fmt"

func digitN(k, n int) int{
    for i := 1; i <= n-1; i++ { 
    k/=10
  }
    if k!=0{
    return k%10
  } else { 
    return -1;
  }
}

func main() {
  var k int 
    for i := 1; i <= 5; i ++{
        fmt.Print("K = ")
        fmt.Scan(&k)
        for j := 1; j <= 5; j ++ {
        fmt.Print(digitN(k,j), " ")
    }
    fmt.Println()
    }
}