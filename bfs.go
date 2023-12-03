package main

import (
	"github.com/golang-collections/collections/queue"
	"reflect"
)

func BFS(start [3][3]int, target [3][3]int) bool {

	visited := make([]*Grid, 0) // 记录所有访问过的结点
	startGrid := &Grid{Pre: nil, CurState: start, TarState: target, BlankPos: nil}

	q := queue.New()     // 记录待访问节点
	q.Enqueue(startGrid) // 先访问起始节点
	for q.Len() > 0 {
		curGrid := q.Dequeue().(*Grid) // 访问队头
		if reflect.DeepEqual(curGrid.CurState, target) {
			curGrid.displayPath() // 到达目标则打印路径
			return true
		}
		// 队头标记为已访问
		visited = append(visited, curGrid)

		// 生成下一批待访问节点
		nextGrids := next(curGrid)

		// 仅计划访问从未访问过的那些节点
		tidyNextGrids(nextGrids, visited, curGrid, q)

	}
	return false

}

func tidyNextGrids(nextGrids []*Grid,
	visited []*Grid,
	curGrid *Grid,
	queue *queue.Queue) {
	for _, nextGrid := range nextGrids {
		for _, g := range visited {
			nextGrid.Pre = curGrid
			if !reflect.DeepEqual(g.CurState, nextGrid.CurState) {
				// 该后继节点未被访问
				queue.Enqueue(nextGrid)
			}
		}
	}
}

func next(grid *Grid) []*Grid {
	grids := make([]*Grid, 0)
	if up := grid.up(); up != nil {
		grids = append(grids, up)
	}
	if left := grid.left(); left != nil {
		grids = append(grids, left)
	}
	if down := grid.down(); down != nil {
		grids = append(grids, down)
	}
	if right := grid.right(); right != nil {
		grids = append(grids, right)
	}
	return grids

}
