package main

import "fmt"

type array []int

func (arr array) quickSort(left int, right int) {
	if right <= left {
		return
	}

	pivot := arr[(left+right)/2]
	i, j := left, right
	for i <= j {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i > j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}

	arr.quickSort(left, j)
	arr.quickSort(i, right)
}

func main() {
	arr := array{5, 2, 7, 1, 0, 9, 2, 4, 6, 3}
	fmt.Println(arr)
	arr.quickSort(0, len(arr)-1)
	fmt.Println(arr)
}
