package main

import (
	"fmt"
	"os"
)

func readMeze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	if err != nil {
		panic(err)
	}

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, row)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze

}

type Point struct {
	i, j int
}

var dirs = [4]Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func (p Point) add(r Point) Point {
	return Point{p.i + r.i, p.j + r.j}
}

func (p Point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end Point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []Point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			if val, ok := next.at(maze); !ok || val == 1 {
				continue
			}

			if val, ok := next.at(steps); !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMeze("D:\\Dev\\go\\spider\\study\\maze\\maze.txt")

	steps := walk(maze, Point{0, 0}, Point{len(maze) - 1, len(maze[0]) - 1})

	//[0 1 0 5 6 7]
	//[1 2 3 4 0 8]
	//[2 3 0 5 0 9]
	//[3 0 0 0 11 10]
	//[4 5 0 13 12 0]
	//[5 6 0 14 13 14]
	for i := range steps {
		fmt.Println(steps[i])
	}
}
