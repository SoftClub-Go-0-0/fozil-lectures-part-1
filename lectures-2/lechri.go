package main

import "fmt"

func InvDigits(n *int){
	var rev int 
	for i:=*n;i>0;i/=10{
		rev=rev*10+i%10
	}
	*n=rev
}

func main() {
	var N int
	for i := 1; i <= 5; i++ {
		fmt.Scan(&N)
		InvDigits(&N)
		fmt.Println(N)
	}
}