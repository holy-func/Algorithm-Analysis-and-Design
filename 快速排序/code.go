package main

import "fmt"

func quickSort(arr *[]int) *[]int {
	if len(*arr) <= 1 {
		return arr
	}
	base, pos := (*arr)[0], 1
	left, right := &[]int{}, &[]int{}
	for pos < len(*arr) {
		if (*arr)[pos] < base {
			*left = append(*left, (*arr)[pos])
		} else {
			*right = append(*right, (*arr)[pos])
		}
		pos++
	}
	left, right = quickSort(left), quickSort(right)
	ret := []int{}
	ret = append(ret, *left...)
	ret = append(ret, base)
	ret = append(ret, *right...)
	return &ret
}
func quickSortEffect(arr *[]int, l, r int) {
	if len(*arr) <= 1 || l >= r {
		return
	}
	right, left, pos := r, l, l
	base := (*arr)[l]
	for left < right {
		for right > left && (*arr)[right] >= base {
			right--
		}
		for left < right && (*arr)[left] <= base {
			left++
		}
		(*arr)[left], (*arr)[right] = (*arr)[right], (*arr)[left]
	}
	(*arr)[pos], (*arr)[left] = (*arr)[left], base
	quickSortEffect(arr, l, left-1)
	quickSortEffect(arr, left+1, r)
}
func main() {
	arr := &[]int{1, 39, 29, 23, 22, 90, 88, 77, 1, 2}
	arr_ := quickSort(arr)
	fmt.Println(arr)
	quickSortEffect(arr, 0, len(*arr)-1)
	fmt.Println(arr)
	fmt.Println(arr_)
	// &[1 39 29 23 22 90 88 77 1 2]
	// &[1 1 2 22 23 29 39 77 88 90]
	// &[1 1 2 22 23 29 39 77 88 90]
}
