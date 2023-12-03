package main

type Point struct {
	Y int // i
	X int // j
}

// Grid 八数码棋盘
type Grid struct {
	Pre        *Grid // 前驱节点
	CurState   [3][3]int
	TarState   [3][3]int
	BlankPos   *Point
	LastAction int // 通过什么操作得到，用于优化求解
}

func (grid *Grid) displayPath() {
	res := make([]*Grid, 0)
	res = append(res, grid)

	cur := grid.Pre
	for cur != nil {
		res = append(res, cur)
		cur = cur.Pre
	}

	// 切片反转
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	for _, elem := range res {
		elem.displayCur()
	}
}

func (grid *Grid) displayCur() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(grid.CurState[i][j], "   ")
		}
		println()
	}
	println("----------------")
}

func (grid *Grid) setBlank() {
	y, x := grid.findPos(0)
	grid.BlankPos = &Point{Y: y, X: x}
}
func (grid *Grid) findPos(num int) (y int, x int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid.CurState[i][j] == num {
				y = i
				x = j
			}
		}
	}
	return
}

func (grid *Grid) move(dy int, dx int) *Grid {
	if grid.BlankPos == nil {
		grid.setBlank()
	}

	g := &Grid{Pre: nil, CurState: grid.CurState, TarState: grid.TarState}
	y := grid.BlankPos.Y
	x := grid.BlankPos.X

	// 移动空白格位置(移动相当于交换)
	g.CurState[y][x], g.CurState[y+dy][x+dx] =
		g.CurState[y+dy][x+dx], g.CurState[y][x]

	// 修改新的空白格位置
	g.BlankPos = &Point{Y: y + dy, X: x + dx}
	return g
}

// 白色（空）方格向上移动
func (grid *Grid) up() *Grid {
	if grid.BlankPos == nil {
		grid.setBlank()
	}
	if grid.BlankPos.Y == 0 {
		return nil
	}
	return grid.move(-1, 0)
}

// 白色（空）方格向下移动
func (grid *Grid) down() *Grid {
	if grid.BlankPos == nil {
		grid.setBlank()
	}
	if grid.BlankPos.Y == 2 {
		return nil
	}
	return grid.move(1, 0)
}

// 白色（空）方格向左移动
func (grid *Grid) left() *Grid {
	if grid.BlankPos == nil {
		grid.setBlank()
	}
	if grid.BlankPos.X == 0 {
		return nil
	}
	return grid.move(0, -1)
}

// 白色（空）方格向右移动
func (grid *Grid) right() *Grid {
	if grid.BlankPos == nil {
		grid.setBlank()
	}
	if grid.BlankPos.X == 2 {
		return nil
	}
	return grid.move(0, 1)
}
