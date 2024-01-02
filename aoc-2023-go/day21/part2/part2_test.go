package part2

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strconv"
	"testing"
)

type Point struct {
	row, col int
	visited  bool
	dist     int
}

var points = make(map[string]*Point)

func key(row, col int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(col)
}

func TestInput(t *testing.T) {
	values := [3]int{result(2), result(4), result(6)}

	m := util.Matrix([][]float64{
		{4, 2, 1},
		{16, 4, 1},
		{36, 6, 1},
	})
	v := util.NewVector(values[0], values[1], values[2])
	_, vSolved := util.Solve(m, v)

	fmt.Printf("%f", vSolved[0]*202300*202300+vSolved[1]*202300+vSolved[2])
}

func result(n int) int {
	target := 65 + 131*n
	input := util.Duplicate2D(n*2+1, n*2+1, Input)
	grid := util.ToGrid(input)
	var queue []*Point

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			point := Point{row: row, col: col, dist: 2*(len(grid)+len(grid[0])) + 1}
			points[key(point.row, point.col)] = &point
			if row == (len(grid)-1)/2 && col == (len(grid[0])-1)/2 {
				queue = append(queue, points[key(point.row, point.col)])
				point.dist = 0
			}
		}
	}

	for i := 0; i < len(queue); i++ {
		iterateQueue := util.Duplicate(queue)
		queue = []*Point{}
		for _, point := range iterateQueue {
			if grid[point.row][point.col] == "#" {
				continue
			}
			if point.visited {
				continue
			}
			point.visited = true
			if point.dist == target {
				continue
			}
			neighbours := []*Point{
				points[key(point.row-1, point.col)],
				points[key(point.row+1, point.col)],
				points[key(point.row, point.col-1)],
				points[key(point.row, point.col+1)],
			}

			for _, neighbour := range neighbours {
				if neighbour == nil {
					continue
				}
				if grid[neighbour.row][neighbour.col] == "#" {
					continue
				}
				if neighbour.dist != target {
					neighbour.dist = point.dist + 1
				}
				queue = append(queue, neighbour)
			}

			if (point.dist+1)%2 == 0 && point.dist == i {
				point.dist = target
			}
		}
	}

	count := 0
	for _, value := range points {
		if value.visited && value.dist == target {
			count++
		}
	}

	return count
}
