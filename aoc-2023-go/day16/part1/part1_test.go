package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"testing"
)

type Point struct {
	row int
	col int
}

type Beam struct {
	pos Point
	dir Point
}

func TestInput(t *testing.T) {
	beams := []Beam{{pos: Point{col: 0, row: 0}, dir: Point{row: 0, col: 1}}}
	visited := make(map[Point]bool)
	grid := util.ToGrid(Input)
	var newBeams []Beam
	silentRounds := 0

	for silentRounds < 100 {
		beams = append(beams, newBeams...)
		beams = removeDuplicates(beams)
		visitedAtStart := len(visited)
		newBeams = []Beam{}
		for i := 0; i < len(beams); i++ {
			beam := &beams[i]
			visited[beam.pos] = true
			if at(beam.pos, grid) == "/" {
				if beam.dir.col != 0 {
					beam.dir.row = -beam.dir.col
					beam.dir.col = 0
				} else {
					beam.dir.col = -beam.dir.row
					beam.dir.row = 0
				}
			}
			if at(beam.pos, grid) == "\\" {
				if beam.dir.col != 0 {
					beam.dir.row = beam.dir.col
					beam.dir.col = 0
				} else {
					beam.dir.col = beam.dir.row
					beam.dir.row = 0
				}
			}
			if at(beam.pos, grid) == "-" && beam.dir.row != 0 {
				beam.dir.row = 0
				beam.dir.col = 1
				newBeam := Beam{pos: beam.pos, dir: beam.dir}
				newBeam.dir.col = -1
				if move(&newBeam, grid) {
					newBeams = append(newBeams, newBeam)
				}
			}
			if at(beam.pos, grid) == "|" && beam.dir.col != 0 {
				beam.dir.row = 1
				beam.dir.col = 0
				newBeam := Beam{pos: beam.pos, dir: beam.dir}
				newBeam.dir.row = -1
				if move(&newBeam, grid) {
					newBeams = append(newBeams, newBeam)
				}
			}
			if !move(beam, grid) {
				beams = append(beams[:i], beams[i+1:]...)
			}
		}

		if visitedAtStart == len(visited) {
			silentRounds++
		} else {
			silentRounds = 0
		}
	}

	fmt.Println(len(visited))
}

func move(beam *Beam, grid [][]string) bool {
	beam.pos.col += beam.dir.col
	beam.pos.row += beam.dir.row

	return util.InGrid(beam.pos.row, beam.pos.col, grid)
}

func at(pos Point, grid [][]string) string {
	return grid[pos.row][pos.col]
}

func removeDuplicates(beams []Beam) []Beam {
	seen := make(map[string]bool)
	var result []Beam

	for _, beam := range beams {
		key := fmt.Sprintf("%d,%d,%d,%d", beam.pos.row, beam.pos.col, beam.dir.col, beam.dir.row)
		if _, exists := seen[key]; !exists {
			seen[key] = true
			result = append(result, beam)
		}
	}

	return result
}
