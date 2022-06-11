package main

import "fmt"

func MatrixChain(rows *[]int) {
	n := len(*rows)
	m, s := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}
	for r := 2; r < n; r++ {
		for i := 1; i < n-r+1; i++ {
			j := i + r - 1
			m[i][j] = m[i+1][j] + (*rows)[i-1]*(*rows)[i]*(*rows)[j]
			s[i][j] = i
			for k := i + 1; k < j; k++ {
				t := m[i][k] + m[k+1][j] + (*rows)[i-1]*(*rows)[k]*(*rows)[j]
				if t < m[i][j] {
					m[i][j] = t
					s[i][j] = k
				}
			}
		}
	}
	for _, v := range m {
		fmt.Println(v)
	}
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println()
}
func Recur(rows *[]int, i, j int) int {
	if i == j {
		return 0
	} else {
		u := Recur(rows, i, i) + Recur(rows, i+1, j) + (*rows)[i]*(*rows)[i+1]*(*rows)[j+1]
		for t := i + 1; t < j; t++ {
			_u := Recur(rows, i, t) + Recur(rows, t+1, j) + (*rows)[i]*(*rows)[t+1]*(*rows)[j+1]
			if _u < u {
				u = _u
			}
		}
		return u
	}
}
func main() {
	MatrixChain(&[]int{30, 35, 15, 5, 10, 20, 25})
	MatrixChain(&[]int{2, 3, 4, 5, 6, 8, 5})
	fmt.Println(Recur(&[]int{30, 35, 15, 5, 10, 20, 25}, 0, 5))
}
