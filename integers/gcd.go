package main

import "fmt"

// Compute gcd(a,b) using recursion
func gcdRecursive(a int, b int) int {
	if b == 0 {
		return a
	}

	return gcdRecursive(b, a%b)
}

// Compute gcd(a,b) using iteration
func gcdIterative(a int, b int) int {
	for b > 0 {
		a, b = b, a%b
	}

	return a
}

// Compute gcd(a,b) using iterative method with subtraction
func gcdSubtractive(a int, b int) int {
	for a > 0 && b > 0 {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}

	if a > 0 {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(gcdRecursive(20, 15))
	fmt.Println(gcdIterative(20, 15))
	fmt.Println(gcdSubtractive(20, 15))
}
