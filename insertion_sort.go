package main

import "fmt"

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j -= 1
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int = []int{1, -10, 4, 2, 1, 10, 3}

	sort(arr)
	fmt.Printf("array: %v\n", arr[:])
}
