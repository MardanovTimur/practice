package main

import "fmt"

func getPointer(arr []int, a, b int) int {
	// Lomuto division
	pointerValue := arr[b]
	var i = a
	for j := a; j < b; j++ {
		if arr[j] <= pointerValue {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}
	arr[i], arr[b] = arr[b], arr[i]
	return i
}

func quickSortRect(arr []int, a, b int) []int {
	if a < b {
		q := getPointer(arr, a, b)
		arr = quickSortRect(arr, a, q-1)
		arr = quickSortRect(arr, q+1, b)
	}
	return arr
}

func quickSort(arr []int) []int {
	return quickSortRect(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{3, 10, 2, -1, 20, 3, -5}
	arr = quickSort(arr)
	fmt.Printf("Sorted arr: %v\n", arr)
}
