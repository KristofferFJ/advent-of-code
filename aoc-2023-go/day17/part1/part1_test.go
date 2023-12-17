package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

type Point struct {
	row, col int
}

type Path struct {
	pos, dir Point
	value    int
	dirMoves int
}

func key(point, dir Point, moves int) string {
	return fmt.Sprintf("%d,%d,%d,%d,%d", point.row, point.col, dir.row, dir.col, moves)
}

func (point Point) at(grid [][]int) int {
	return grid[point.row][point.col]
}

func TestInput(t *testing.T) {
	grid := util.ToIntGrid(Input)
	bestPaths := make(map[string]int)
	paths := []Path{
		{pos: Point{1, 0}, dir: Point{1, 0}, value: 0, dirMoves: 1},
		{pos: Point{0, 1}, dir: Point{0, 1}, value: 0, dirMoves: 1},
	}
	minScore := 1000000000000000000

	for len(paths) > 0 {
		path := &paths[0]
		path.value += path.pos.at(grid)
		key := key(path.pos, path.dir, path.dirMoves)
		_, keyExists := bestPaths[key]
		if path.value > minScore || (keyExists && bestPaths[key] <= path.value) {
			paths = paths[1:]
			continue
		} else {
			bestPaths[key] = path.value
		}

		if path.pos.row == len(grid)-1 && path.pos.col == len(grid[0])-1 {
			paths = paths[1:]
			if path.value < minScore {
				minScore = path.value
				fmt.Printf("New min score:%d\n", minScore)
			}
			continue
		}
		paths = paths[1:]
		paths = append(paths, getAdjacent(*path, grid)...)
	}

	fmt.Println(minScore)
}

func getAdjacent(path Path, grid [][]int) []Path {
	var adjacentPaths []Path
	directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for _, dir := range directions {
		newPos := Point{path.pos.row + dir.row, path.pos.col + dir.col}
		if !util.InGrid(newPos.row, newPos.col, grid) {
			continue
		}
		if path.dir.row*dir.row+path.dir.col*dir.col == -1 {
			continue
		}
		sameDir := dir.row == path.dir.row && dir.col == path.dir.col
		dirMoves := 0
		if sameDir {
			if path.dirMoves == 2 {
				continue
			}
			dirMoves = path.dirMoves + 1
		}
		adjacentPaths = append(adjacentPaths, Path{newPos, dir, path.value, dirMoves})
	}

	return adjacentPaths
}
