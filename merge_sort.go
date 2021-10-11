package main

import "fmt"

var MAX_INT int = int(1) << 32

func MergeFunc(arr []int, a, q, b int) []int {
	var L, R = make([]int, q-a+2), make([]int, b-q+2)
	var i, j int
	for i = a; i <= q; i++ {
		L[i-a] = arr[i]
	}
	L[i-a] = MAX_INT

	for i = q + 1; i <= b; i++ {
		R[i-q-1] = arr[i]
	}
	R[i-q-1] = MAX_INT

	i, j = 0, 0
	for k := a; k <= b; k++ {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i += 1
		} else {
			arr[k] = R[j]
			j += 1
		}
	}
	return arr
}

func MergeSort(arr []int, a, b int) []int {
	var q int
	if a < b {
		q = (a + b) / 2
		MergeSort(arr, a, q)
		MergeSort(arr, q+1, b)
		arr = MergeFunc(arr, a, q, b)
	}
	return arr
}

func main() {
	arr := MergeSort([]int{1, -10, 2, -1, 20, 30, 5, 2, 7}, 0, 8)
	fmt.Printf("Sorted: %v\n", arr)
}
