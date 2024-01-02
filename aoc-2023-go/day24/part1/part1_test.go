package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

var InputTest = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

type Point struct{ x, y, z float64 }
type Hail struct{ pos, vel Point }
type Trajectory struct{ a, b float64 }

var zero = float64(0)

func TestInput(t *testing.T) {
	hails := parseInput(Input)

	minVal := float64(200000000000000)
	maxVal := float64(400000000000000)

	count := 0
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			hail1 := hails[i]
			hail2 := hails[j]
			x, y, t1, t2 := getCrossing(hail1, hail2)
			if x >= minVal &&
				x <= maxVal &&
				y >= minVal &&
				y <= maxVal &&
				t1 > zero &&
				t2 > zero {
				count++
			}
		}
	}

	fmt.Println(count)
}

func getTrajectory(h Hail) Trajectory {
	return Trajectory{
		a: h.vel.y / (h.vel.x),
		b: -h.vel.y*(h.pos.x/(h.vel.x)) + h.pos.y,
	}
}

func getCrossing(h1, h2 Hail) (float64, float64, float64, float64) {
	trajectory1 := getTrajectory(h1)
	trajectory2 := getTrajectory(h2)
	if trajectory1.a == trajectory2.a {
		return 0, 0, 0, 0
	}
	x := (trajectory2.b - trajectory1.b) / (trajectory1.a - trajectory2.a)
	y := (trajectory1.a * x) + trajectory1.b
	t1 := (x - h1.pos.x) / h1.vel.x
	t2 := (x - h2.pos.x) / h2.vel.x
	return x, y, t1, t2
}

func parseInput(input string) (hails []Hail) {
	for _, line := range strings.Split(input, "\n") {
		split := util.IntArray(line)
		hails = append(hails,
			Hail{
				pos: Point{float64(split[0]), float64(split[1]), float64(split[2])},
				vel: Point{float64(split[3]), float64(split[4]), float64(split[5])},
			})
	}

	return hails
}
