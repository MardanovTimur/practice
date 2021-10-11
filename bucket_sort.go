package main

import "fmt"

var bucketSize = 401
var bucket []int = make([]int, bucketSize)

func bucketSort(arr []int) []int {
	for _, v := range arr {
		bucket[200+v] += 1
	}

	var x int = 0
	for i := 0; i <= 400; i++ {
		for j := 0; j < bucket[i]; j++ {
			arr[x] = i - 200
			x += 1
		}
	}
	return arr
}

func main() {
	//  -200 <= a[n] <= 200
	arr := []int{23, 31, 1, 8, 2, -10, 3, 20, 100, -20, -10}
	arr = bucketSort(arr)
	fmt.Printf("%v\n", arr)
}
