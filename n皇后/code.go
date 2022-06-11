package main

import "fmt"

func Queen(n int) *[][]int {
	x, used, Xs := make([]int, n), make([]bool, n), [][]int{}
	var backtrack func(int)
	isXWorkable := func(cur int) bool {
		j := cur - 1
		for i := 0; i < cur-1; i++ {
			// for j := i + 1; j < cur; j++ {
			if i+x[i] == j+x[j] || x[j]-x[i] == j-i {
				return false
			}
			// }
		}
		return true
	}
	backtrack = func(cur int) {
		if !isXWorkable(cur) {
			return
		}
		if cur == n {
			X := make([]int, n)
			copy(X, x)
			Xs = append(Xs, X)
			return
		}
		for i := 0; i < n; i++ {
			if !used[i] {
				used[i] = true
				x[cur] = i
				backtrack(cur + 1)
				used[i] = false
			}
		}
	}
	backtrack(0)
	return &Xs
}
func main() {
	Xs := Queen(5)
	fmt.Println(len(*Xs))
	for _, v := range *Xs {
		fmt.Println(v)
	}
	// 10
	// [0 2 4 1 3]
	// [0 3 1 4 2]
	// [1 3 0 2 4]
	// [1 4 2 0 3]
	// [2 0 3 1 4]
	// [2 4 1 3 0]
	// [3 0 2 4 1]
	// [3 1 4 2 0]
	// [4 1 3 0 2]
	// [4 2 0 3 1]
}
