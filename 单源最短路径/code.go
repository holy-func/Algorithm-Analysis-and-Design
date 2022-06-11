package main

import (
	"fmt"
	"sort"
)

type key interface {
	~int | ~string
}
type node[T key] struct {
	origin T
	to     T
	dist   int
}

func min(a, b int) int {
	if (a != -1 && a < b) || b == -1 {
		return a
	} else {
		return b
	}
}
func Dijkstra[T key](g *map[T]map[T]int, origin T) *map[T]int {
	dist := map[T]int{}
	S, U := map[T]node[T]{origin: {origin, origin, 0}}, []*node[T]{}
	settled := map[T]bool{origin: true}
	for k, v := range (*g)[origin] {
		if k != origin {
			U = append(U, &node[T]{origin, k, v})
			settled[k] = true
		}
	}

	for len(U) > 0 {
		sort.Slice(U, func(i, j int) bool {
			return min(U[i].dist, U[j].dist) == U[i].dist
		})
		cur := U[0]
		U = U[1:]
		for k := range (*g)[cur.to] {
			if _, ok := settled[k]; !ok {
				U = append(U, &node[T]{origin, k, -1})
				settled[k] = true
			}
		}
		S[cur.to] = *cur
		for _, u := range U {
			if dist, connected := (*g)[cur.to][u.to]; connected {
				u.dist = min(u.dist, cur.dist+dist)
			}
		}
	}
	for _, s := range S {
		dist[s.to] = s.dist
	}
	return &dist
}
func main() {
	ret := Dijkstra(&map[int]map[int]int{
		0: {0: 0, 1: 20, 3: 40},
		1: {2: 5, 3: 4},
		2: {4: 10, 3: 9},
		3: {9: 1, 8: 7},
	}, 0)
	fmt.Println(ret)
	// &map[0:0 1:20 2:25 3:24 4:35 8:31 9:25]
}
