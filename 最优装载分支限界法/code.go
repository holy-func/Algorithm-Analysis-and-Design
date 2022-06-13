package main

import (
	"fmt"
	"github.com/holy-func/heap"
)

type node struct {
	i        int
	state    bool
	priority int
	w        int
	parent   *node
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func MaxLoading(goods []int, c1, c2 int) bool {
	w, bestw := 0, 0
	for _, good := range goods {
		w += good
	}
	priorityQueue := heap.MaxHeap(func(a, b *node) bool { return a.priority < b.priority })
	priorityQueue.Push(&node{0, false, w, 0, nil})
	for {
		cur := priorityQueue.Pop()
		if cur.i > 0 {
			fmt.Println(fmt.Sprintf("第%d个货物,是否选择:%v,当前节点优先级:%d,当前装载量:%d,当前最优装载量:%d", cur.i, cur.state, cur.priority, cur.w, bestw))
		}
		if cur.i < len(goods) {
			good := goods[cur.i]
			if priority := cur.priority - good; priority >= bestw {
				priorityQueue.Push(&node{cur.i + 1, false, priority, cur.w, cur})
			}
			if good+cur.w <= c1 {
				curw := cur.w + good
				priorityQueue.Push(&node{cur.i + 1, true, cur.priority, curw, cur})
				bestw = max(bestw, curw)
			}
		} else {
			choose := make([]bool, len(goods))
			for cur != nil && cur.i > 0 {
				choose[cur.i-1] = cur.state
				cur = cur.parent
			}
			fmt.Println("装载车1的最优装载选择为:", choose)
			fmt.Println(fmt.Sprintf("总装载量:%d,装载车1装载量为:%d,装载车2可装载量为:%d,装载车2需装载量为:%d", w, bestw, c2, w-bestw))
			break
		}
	}
	return w-bestw <= c2
}
func main() {
	success := MaxLoading([]int{10, 30, 50, 5, 5, 8}, 68, 40)
	fmt.Println(success)
	// 第1个货物,是否选择:true,当前节点优先级:108,当前装载量:10,当前最优装载量:10
	// 第2个货物,是否选择:true,当前节点优先级:108,当前装载量:40,当前最优装载量:40
	// 第1个货物,是否选择:false,当前节点优先级:98,当前装载量:0,当前最优装载量:40
	// 第2个货物,是否选择:true,当前节点优先级:98,当前装载量:30,当前最优装载量:40
	// 第2个货物,是否选择:false,当前节点优先级:78,当前装载量:10,当前最优装载量:40
	// 第3个货物,是否选择:true,当前节点优先级:78,当前装载量:60,当前最优装载量:60
	// 第4个货物,是否选择:true,当前节点优先级:78,当前装载量:65,当前最优装载量:65
	// 第5个货物,是否选择:false,当前节点优先级:73,当前装载量:65,当前最优装载量:65
	// 第4个货物,是否选择:false,当前节点优先级:73,当前装载量:60,当前最优装载量:65
	// 第5个货物,是否选择:true,当前节点优先级:73,当前装载量:65,当前最优装载量:65
	// 第5个货物,是否选择:false,当前节点优先级:68,当前装载量:60,当前最优装载量:65
	// 第6个货物,是否选择:true,当前节点优先级:68,当前装载量:68,当前最优装载量:68
	// 装载车1的最优装载选择为: [true false true false false true]
	// 总装载量:108,装载车1装载量为:68,装载车2可装载量为:40,装载车2需装载量为:40
	// true
}
