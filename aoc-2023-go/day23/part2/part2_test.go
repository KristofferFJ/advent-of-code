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

type SubPath struct {
	from, to Point
	dist     int
}

type UnfinishedSubPath struct {
	start, previous, pos Point
	dist                 int
}

type Next struct {
	to  Point
	len int
}

type Path struct {
	pos  Point
	path []Point
	len  int
}

var SubPaths = make(map[Point][]Next)

var visitedIntersection = make(map[Point]bool)

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	end := Point{row: len(grid) - 1, col: len(grid[0]) - 2}
	start := Point{row: 0, col: 1}
	subPaths := findSubPaths(start, grid, end)
	for _, subPath := range subPaths {
		target, ok := SubPaths[subPath.from]
		if ok {
			SubPaths[subPath.from] = append(target, Next{subPath.to, subPath.dist})
		} else {
			for _, value := range SubPaths[subPath.from] {
				if value.to == subPath.to {
					continue
				}
			}
			SubPaths[subPath.from] = []Next{{subPath.to, subPath.dist}}
		}
		target, ok = SubPaths[subPath.to]
		if ok {
			SubPaths[subPath.to] = append(target, Next{subPath.from, subPath.dist})
		} else {
			for _, value := range SubPaths[subPath.to] {
				if value.to == subPath.from {
					continue
				}
			}
			SubPaths[subPath.to] = []Next{{subPath.from, subPath.dist}}
		}
	}

	for key, value := range SubPaths {
		SubPaths[key] = util.RemoveDuplicateValue(value)
	}
	startPath := Path{pos: start, len: 0, path: []Point{start}}
	paths := createPaths(startPath, end)

	maxLength := 0
	for _, path := range paths {
		if path.len > maxLength {
			maxLength = path.len
		}
	}
	fmt.Println(maxLength)
}

func createPaths(path Path, end Point) []Path {
	if path.pos == end {
		return []Path{path}
	}
	var paths []Path
	for _, next := range SubPaths[path.pos] {
		if PointInList(next.to, path.path) {
			continue
		}
		newPath := Path{
			pos:  next.to,
			len:  path.len + next.len,
			path: append(path.path, path.pos),
		}
		paths = append(paths, createPaths(newPath, end)...)
	}
	return paths
}

func findSubPaths(start Point, grid [][]string, end Point) (subPaths []SubPath) {
	paths := []UnfinishedSubPath{{start, start, start, 0}}
	for len(paths) > 0 {
		path := paths[len(paths)-1]
		paths = paths[0 : len(paths)-1]

		var newPaths []UnfinishedSubPath
		for _, direction := range []Point{Up, Down, Left, Right} {
			lastPoint := path.pos
			nextPoint := Point{
				row: lastPoint.row + direction.row,
				col: lastPoint.col + direction.col,
			}
			if util.InGrid(nextPoint.col, nextPoint.row, grid) == false {
				continue
			}
			if grid[nextPoint.row][nextPoint.col] == "#" {
				continue
			}
			if nextPoint == path.previous {
				continue
			}
			newPaths = append(newPaths, UnfinishedSubPath{
				start:    path.start,
				previous: path.pos,
				pos:      nextPoint,
				dist:     path.dist + 1,
			})
		}
		if len(newPaths) == 0 {
			continue
		}
		if len(newPaths) == 1 {
			if newPaths[0].pos == end {
				subPaths = append(subPaths, SubPath{
					from: path.start,
					to:   end,
					dist: path.dist + 1,
				})
			}
			paths = append(paths, newPaths[0])
			continue
		}
		subPaths = append(subPaths, SubPath{
			from: path.start,
			to:   path.pos,
			dist: path.dist,
		})
		if !visitedIntersection[path.pos] == true {
			for _, newPath := range newPaths {
				paths = append(paths, UnfinishedSubPath{
					start:    path.pos,
					previous: path.pos,
					pos:      newPath.pos,
					dist:     1,
				})
			}
			visitedIntersection[path.pos] = true
		}
	}

	subPaths = removeDuplicates(subPaths)

	for _, sub := range subPaths {
		fmt.Printf("Subpath %d,%d to %d,%d, dist %d\n", sub.from.row, sub.from.col, sub.to.row, sub.to.col, sub.dist)
	}

	return subPaths
}

func removeDuplicates(subPaths []SubPath) []SubPath {
	seen := make(map[string]bool)
	var result []SubPath

	for _, subPath := range subPaths {
		key := fmt.Sprintf("%d,%d,%d,%d", subPath.from.row, subPath.from.col, subPath.to.col, subPath.to.row)
		if _, exists := seen[key]; !exists {
			seen[key] = true
			result = append(result, subPath)
		}
	}

	return result
}

func PointInList(point Point, list []Point) bool {
	for _, elem := range list {
		if elem.row == point.row && elem.col == point.col {
			return true
		}
	}
	return false
}
