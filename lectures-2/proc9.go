package main

import "fmt"

func AddLeftDigit(d int, k *int) {
	for i := *k; i > 0; i /= 10 {
		d *= 10
	}
	v := d + *k
	*k = v
}

func main() {
	var K, D1, D2 int
	fmt.Scan(&K, &D1, &D2)
	for i := 1; i <= 2; i++ {
		if i == 1 {
			AddLeftDigit(D1, &K)
			fmt.Println(K)

		} else {
			AddLeftDigit(D2, &K)
			fmt.Println(K)

		}
	}
}
