package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

var InputTest = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

type Point struct {
	row, col int
}

var (
	Up    = Point{row: -1, col: 0}
	Down  = Point{row: 1, col: 0}
	Left  = Point{row: 0, col: -1}
	Right = Point{row: 0, col: 1}
)

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	end := Point{row: len(grid) - 1, col: len(grid[0]) - 2}
	start := Point{row: 0, col: 1}
	maxLength := 0
	paths := [][]Point{{start}}
	for len(paths) > 0 {
		path := paths[0]
		paths = paths[1:]
		if path[len(path)-1] == end {
			if len(path) > maxLength {
				maxLength = len(path)
			}
			continue
		}

		for _, direction := range []Point{Up, Down, Left, Right} {
			lastPoint := path[len(path)-1]
			nextPoint := Point{
				row: lastPoint.row + direction.row,
				col: lastPoint.col + direction.col,
			}
			if util.InGrid(nextPoint.col, nextPoint.row, grid) == false {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == "<" && direction != Left {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == ">" && direction != Right {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == "^" && direction != Up {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == "v" && direction != Down {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == "#" {
				continue
			}
			if util.Contains(path, nextPoint) {
				continue
			}
			newPath := util.Duplicate(path)
			newPath = append(newPath, nextPoint)
			paths = append(paths, newPath)
		}
	}

	fmt.Println(maxLength - 1)
}
