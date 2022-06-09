package main

import "fmt"

func binarySearch(arr *[]int, target int) int {
	left, right := 0, len(*arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		v := (*arr)[mid]
		if v < target {
			left = mid + 1
		} else if v > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
func main() {
	arr := []int{1, 2, 3, 3, 4, 6, 30, 90, 118}
	fmt.Println(binarySearch(&arr, 2), binarySearch(&arr, 6), binarySearch(&arr, 118))
	// 1 5 8
}
