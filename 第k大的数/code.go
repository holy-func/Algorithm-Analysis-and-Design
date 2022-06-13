package main

import "fmt"

type num interface {
	~int | ~float32
}

func FindKthLargest[T num](arr *[]T, k int) T {
	if k > len(*arr) {
		panic("error K")
	}
	low, high := 0, len(*arr)-1
	var cur int
	for low < high {
		cur = QuickSelect(arr, low, high)
		if cur > k-1 {
			high = cur - 1
		} else if cur < k-1 {
			low = cur + 1
		} else {
			return (*arr)[cur]
		}
	}
	return (*arr)[low]
}
func QuickSelect[T num](arr *[]T, l, r int) int {
	i, j, base := l+1, r, (*arr)[l]
	for i < j {
		for i < j && (*arr)[j] <= base {
			j--
		}
		for i < j && (*arr)[i] >= base {
			i++
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
	if (*arr)[i] > (*arr)[l] {
		(*arr)[i], (*arr)[l] = (*arr)[l], (*arr)[i]
		return i
	} else {
		return l
	}
}
func main() {
	arr := &[]int{1, 22, 3, 44, 5, 66}
	for i := 1; i <= len(*arr); i++ {
		fmt.Println("第", i, "大:", FindKthLargest(arr, i))
	}
	// 第 1 大: 66
	// 第 2 大: 44
	// 第 3 大: 22
	// 第 4 大: 5
	// 第 5 大: 3
	// 第 6 大: 1
}
