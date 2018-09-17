package main

import "fmt"

type array []int

func (arr array) bubbleSort() {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := array{5, 2, 7, 1, 0, 9, 2, 4, 6, 3}
	fmt.Println(arr)
	arr.bubbleSort()
	fmt.Println(arr)
}
