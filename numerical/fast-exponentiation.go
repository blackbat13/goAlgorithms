package main

import "fmt"

func fastExp(a int, n int) int {
	w := 1
	for n > 0 {
		if n & 1 == 1 {
			w *= a
		}

		a *= a
		n /= 2
	}

	return w
}

func main() {
	fmt.Println(fastExp(2, 10))
}
