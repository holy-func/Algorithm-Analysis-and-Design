package main

import "fmt"

type good[T any] struct {
	id T
	w  int
	p  int
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func Knapsack[T any](c int, goods *[]*good[T]) (int, []T) {
	n := len(*goods)
	dp := make([][]int, c+1)
	for i := 0; i <= c; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		cur := (*goods)[i-1]
		for v := 1; v <= c; v++ {
			if v < cur.w {
				dp[v][i] = dp[v][i-1]
				continue
			}
			dp[v][i] = max(dp[v][i-1], dp[v-cur.w][i-1]+cur.p)
		}
	}
	choose := []T{}
	var Traceback func(int, int)
	Traceback = func(i, v int) {
		if i == 0 || v == 0 {
			return
		}
		if dp[v][i] == dp[v][i-1] {
			Traceback(i-1, v)
		} else {
			cur := (*goods)[i-1]
			choose = append(choose, cur.id)
			Traceback(i-1, v-cur.w)
		}
	}
	Traceback(n, c)
	for _, v := range dp {
		fmt.Println(v)
	}
	return dp[c][n], choose
}
func main() {
	best, choose := Knapsack(8, &[]*good[string]{
		{"苹果", 2, 3},
		{"番茄", 3, 4},
		{"玉米", 4, 5},
		{"土豆", 5, 6},
	})
	fmt.Println(best, choose)
	// [0 0 0 0 0]
	// [0 0 0 0 0]
	// [0 3 3 3 3]
	// [0 3 4 4 4]
	// [0 3 4 5 5]
	// [0 3 7 7 7]
	// [0 3 7 8 8]
	// [0 3 7 9 9]
	// [0 3 7 9 10]
	// 10 [土豆 番茄]
}
