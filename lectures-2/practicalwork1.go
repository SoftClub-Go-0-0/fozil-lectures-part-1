package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n2:=n*2
	fmt.Println("It is",n2/60,"hours",n2%60,"minutes.")
}