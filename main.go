package main

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

	BFS(start, target)
}
