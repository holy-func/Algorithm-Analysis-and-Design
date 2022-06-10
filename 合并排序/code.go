package main

import "fmt"

func mergeSort(arr *[]int) *[]int {
	if len(*arr) == 1 {
		return arr
	} else {
		mid := len(*arr) >> 1
		left, right := (*arr)[:mid], (*arr)[mid:]
		left, right = *mergeSort(&left), *mergeSort(&right)
		ret := []int{}
		l, r := 0, 0
		for len(left) > l && len(right) > r {
			if left[l] < right[r] {
				ret = append(ret, left[l])
				l++
			} else {
				ret = append(ret, right[r])
				r++
			}
		}
		if l < len(left) {
			ret = append(ret, left[l:]...)
		}
		if r < len(right) {
			ret = append(ret, right[r:]...)
		}
		return &ret
	}
}
func main() {
	arr := &[]int{1, 39, 29, 11, 22, 90, 88, 77, 1, 2}
	fmt.Println(arr)
	arr = mergeSort(arr)
	fmt.Println(arr)
	// &[1 39 29 11 22 90 88 77 1 2]
	// &[1 1 2 11 22 29 39 77 88 90]
}
