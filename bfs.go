package main

import (
	"github.com/golang-collections/collections/queue"
	"reflect"
)

func BFS(start [3][3]int, target [3][3]int) bool {

	visited := make([]*Grid, 0)
	startGrid := &Grid{Pre: nil, CurState: start, TarState: target, BlankPos: nil}
	q := queue.New()
	q.Enqueue(startGrid)
	for q.Len() > 0 {
		curGrid := q.Dequeue().(*Grid)
		if reflect.DeepEqual(curGrid.CurState, target) {
			curGrid.displayPath()
			return true
		}
		visited = append(visited, curGrid)
		nextGrids := next(curGrid)
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
