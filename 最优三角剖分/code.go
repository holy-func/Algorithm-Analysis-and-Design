package main

import "fmt"

type lineType interface {
	id() int
	value() int
}
type line struct {
	_id    int
	_value int
}

func (l *line) id() int {
	return l._id
}
func (l *line) value() int {
	return l._value
}

type W func(lineType, lineType, lineType) int

// 权值函数
func Triangulation(p *[]lineType, w W) {
	n := len(*p)
	m, s := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}
	for r := 2; r < n; r++ {
		for i := 1; i < n-r+1; i++ {
			j := i + r - 1
			m[i][j] = m[i+1][j] + w((*p)[i-1], (*p)[i], (*p)[j])
			s[i][j] = (*p)[i-1].id()
			for k := i + 1; k < j; k++ {
				t := m[i][k] + m[k+1][j] + w((*p)[i-1], (*p)[k], (*p)[j])
				if t < m[i][j] {
					m[i][j] = t
					s[i][j] = (*p)[k-1].id()
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
func main() {
	Triangulation(&[]lineType{
		&line{9, 30},
		&line{8, 35},
		&line{7, 15},
		&line{6, 5},
		&line{5, 10},
		&line{4, 20},
		&line{3, 25},
	}, func(lt1, lt2, lt3 lineType) int {
		return lt1.value() * lt2.value() * lt3.value()
	})
	// 此权值函数等价于矩阵连乘
}
