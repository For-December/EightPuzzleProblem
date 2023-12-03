package main

import "fmt"

func main() {

	start := [3][3]int{
		{2, 8, 3},
		{1, 0, 4},
		{7, 6, 5},
	}
	target := [3][3]int{
		{1, 2, 3},
		{8, 0, 4},
		{7, 6, 5},
	}
	println("start -> target")
	DisplayData(start, target)

	println("######## =BFS= ########")
	BFS(start, target)

	println("######## =AStar= ########")
	AStar(start, target)
}

func DisplayData(start [3][3]int, target [3][3]int) {
	for i := 0; i < 3; i++ {
		startStr := ""
		targetStr := ""
		for j := 0; j < 3; j++ {
			startStr += fmt.Sprintf("%d  ", start[i][j])
			targetStr += fmt.Sprintf("%d  ", target[i][j])
		}
		if i == 1 {
			println(startStr + "    =>    " + targetStr)
		} else {
			println(startStr + "          " + targetStr)
		}
	}

}
