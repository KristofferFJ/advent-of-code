package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"slices"
	"strconv"
	"testing"
)

var InputTest = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

type Point struct {
	row, col int
	visited  bool
	dist     int
}

var points = make(map[string]*Point)

func key(row, col int) string {
	return strconv.Itoa(row) + "," + strconv.Itoa(col)
}

const TARGET = 64

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	var queue []*Point

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			point := Point{row: row, col: col, dist: 2*(len(grid)+len(grid[0])) + 1}
			points[key(point.row, point.col)] = &point
			if grid[row][col] == "S" {
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
			if point.dist == TARGET {
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
				if neighbour.dist != TARGET {
					neighbour.dist = point.dist + 1
				}
				queue = append(queue, neighbour)
			}

			if point.dist%2 == 0 && point.dist == i {
				point.dist = TARGET
			}
		}
	}

	var count []string
	for _, value := range points {
		if value.visited && value.dist == TARGET {
			count = append(count, key(value.row, value.col))
		}
	}
	slices.Sort(count)
	fmt.Println(len(count))
}
