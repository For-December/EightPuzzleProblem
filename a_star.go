package main

import (
	"reflect"
	"sort"
)

func AStar(start [3][3]int, target [3][3]int) bool {

	openTable := make([]*Grid, 0)  // 存储待考虑的节点
	closeTable := make([]*Grid, 0) // 存储已考虑过的节点

	startGrid := &Grid{Pre: nil, CurState: start, TarState: target, BlankPos: nil}
	updateFn(startGrid)

	// 先访问初始结点
	openTable = append(openTable, startGrid)
	for len(openTable) > 0 {

		// 升序
		sort.Slice(openTable, func(i, j int) bool {
			return openTable[i].Fn.F < openTable[j].Fn.F
		})

		// 选择 open 表中启发函数最小的节点，进行访问，并移动至 close 表
		curGrid := openTable[0]
		openTable = openTable[1:]
		closeTable = append(closeTable, curGrid)

		// 到达目标棋盘则输出
		if reflect.DeepEqual(curGrid.CurState, target) {
			curGrid.displayPath()
			return true
		}

		// 扩展当前结点的子节点
		nextGrids := curGrid.next()
		openTable = tidyNext(nextGrids, openTable, closeTable, curGrid)

	}
	return false
}

func inGrids(elem *Grid, array []*Grid) int {
	for i, grid := range array {
		if reflect.DeepEqual(elem.CurState, grid.CurState) {
			return i
		}
	}
	return -1
}

func tidyNext(nextGrids []*Grid, openTable []*Grid, closeTable []*Grid, curGrid *Grid) []*Grid {
	for _, grid := range nextGrids {
		grid.Pre = curGrid
		updateFn(grid)

		// 如果在 close 表里（之前被访问过），且 close 表里的 F 更大，则更新 close 表
		// 同时将新扩展的结点移入 open 表
		if index := inGrids(grid, closeTable); index != -1 &&
			grid.Fn.F < closeTable[index].Fn.F {
			closeTable = append(closeTable[:index], closeTable[index+1:]...)
			openTable = append(openTable, grid)
			continue
		}

		// 如果在 open 表里（之后会被访问），且 open 表里的 F 更大，则更新 open 表
		if index := inGrids(grid, openTable); index != -1 &&
			grid.Fn.F < openTable[index].Fn.F {
			openTable[index] = grid
			continue
		}

		// 两个表都不在，则直接添加到 open表中
		openTable = append(openTable, grid)
	}
	return openTable

}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func updateFn(grid *Grid) {
	grid.Fn = &Heuristic{}
	// 更新G
	if grid.Pre == nil {
		grid.Fn.G = 0
	} else {
		grid.Fn.G = grid.Pre.Fn.G + 1
	}

	// 更新H
	grid.Fn.H = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// 当前点在目标棋盘中的坐标
			y, x := grid.findPos(grid.TarState[i][j])

			// 当前点回到目标棋盘位置所需绝对步数
			grid.Fn.H += absInt(y-i) + absInt(x-j)
		}
	}

	// 更新F (F=G+H)
	grid.Fn.F = grid.Fn.G + grid.Fn.H

}
