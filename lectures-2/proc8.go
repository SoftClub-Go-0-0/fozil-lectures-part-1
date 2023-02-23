package main

import "fmt"

func AddRightDigit(d, k *int) {
	*k = *k*10 + *d
}

func main() {
	var K, D1, D2 int
	fmt.Scan(&K, &D1, &D2)
	for i := 1; i <= 2; i++ {
		if i == 1 {
			AddRightDigit(&D1, &K)
			fmt.Println(K)

		} else {
			AddRightDigit(&D2, &K)
			fmt.Println(K)

		}
	}
}
