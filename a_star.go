package main

import (
	"github.com/golang-collections/collections/queue"
	"reflect"
)

func AStar(start [3][3]int, target [3][3]int) bool {

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
