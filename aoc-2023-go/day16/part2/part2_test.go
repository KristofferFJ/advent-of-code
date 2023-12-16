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

func (beam Beam) Key() string {
	return fmt.Sprintf("%d,%d,%d,%d", beam.pos.row, beam.pos.col, beam.dir.col, beam.dir.row)
}

func TestInput(t *testing.T) {
	grid := util.ToGrid(Input)
	var startBeams []Beam
	for row := 0; row < len(grid); row++ {
		startBeams = append(startBeams, Beam{pos: Point{row: row, col: 0}, dir: Point{row: 0, col: 1}})
		startBeams = append(startBeams, Beam{pos: Point{row: row, col: len(grid[0]) - 1}, dir: Point{row: 0, col: -1}})
	}
	for col := 0; col < len(grid[0]); col++ {
		startBeams = append(startBeams, Beam{pos: Point{row: 0, col: col}, dir: Point{row: 1, col: 0}})
		startBeams = append(startBeams, Beam{pos: Point{row: len(grid) - 1, col: col}, dir: Point{row: -1, col: 0}})
	}

	countByStartBeam := make(map[Beam]int)
	for _, startBeam := range startBeams {
		beams := []Beam{startBeam}
		visited := make(map[Point]bool)
		var newBeams []Beam
		seenBeam := make(map[string]bool)

		for len(beams) > 0 {
			for i := 0; i < len(beams); i++ {
				if seenBeam[beams[i].Key()] {
					beams = append(beams[:i], beams[i+1:]...)
				}
			}

			beams = append(beams, newBeams...)
			beams = removeDuplicates(beams)
			newBeams = []Beam{}
			for i := 0; i < len(beams); i++ {
				beam := &beams[i]
				seenBeam[beam.Key()] = true
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
		}
		countByStartBeam[startBeam] = len(visited)
	}

	max := 0
	for _, count := range countByStartBeam {
		if count > max {
			max = count
		}
	}
	fmt.Println(max)
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
