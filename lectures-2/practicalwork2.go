package main

import "fmt"

func main() {
	fmt.Printf("%.2f",3.6785)
	fmt.Println()

	fmt.Printf("%T",3.6785)
	fmt.Println()

	fmt.Printf("%s:\t%d", "Decimal", 11)
	fmt.Println()

	fmt.Printf("%s:\t%d", "binary", 11)
	fmt.Println()

	fmt.Printf("%s:\t%o\n", "Octal", 11)

	fmt.Printf("%s:\t%x\n", "Hexadecimal", 11)
}