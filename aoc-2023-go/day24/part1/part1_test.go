package part1

import (
	"fmt"
	"github.com/shopspring/decimal"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

var InputTest = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

type Point struct{ x, y, z decimal.Decimal }
type Hail struct{ pos, vel Point }
type Trajectory struct{ a, b decimal.Decimal }

var zero = decimal.NewFromInt(int64(0))

func TestInput(t *testing.T) {
	hails := parseInput(Input)

	minVal := decimal.NewFromInt(int64(200000000000000))
	maxVal := decimal.NewFromInt(int64(400000000000000))

	count := 0
	for i := 0; i < len(hails)-1; i++ {
		for j := i + 1; j < len(hails); j++ {
			hail1 := hails[i]
			hail2 := hails[j]
			x, y, t1, t2 := getCrossing(hail1, hail2)
			if x.GreaterThanOrEqual(minVal) &&
				x.LessThanOrEqual(maxVal) &&
				y.GreaterThanOrEqual(minVal) &&
				y.LessThanOrEqual(maxVal) &&
				t1.GreaterThan(zero) &&
				t2.GreaterThan(zero) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func getTrajectory(h Hail) Trajectory {
	return Trajectory{
		a: h.vel.y.Div(h.vel.x),
		b: (h.vel.y.Neg().Mul(h.pos.x.Div(h.vel.x))).Add(h.pos.y),
	}
}

func getCrossing(h1, h2 Hail) (decimal.Decimal, decimal.Decimal, decimal.Decimal, decimal.Decimal) {
	trajectory1 := getTrajectory(h1)
	trajectory2 := getTrajectory(h2)
	if trajectory1.a.Equals(trajectory2.a) {
		return decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}, decimal.Decimal{}
	}
	x := (trajectory2.b.Sub(trajectory1.b)).Div(trajectory1.a.Sub(trajectory2.a))
	y := (trajectory1.a.Mul(x)).Add(trajectory1.b)
	t1 := (x.Sub(h1.pos.x)).Div(h1.vel.x)
	t2 := (x.Sub(h2.pos.x)).Div(h2.vel.x)
	return x, y, t1, t2
}

func parseInput(input string) (hails []Hail) {
	for _, line := range strings.Split(input, "\n") {
		split := util.IntArray(line)
		hails = append(hails,
			Hail{
				pos: Point{decimal.NewFromInt(int64(split[0])), decimal.NewFromInt(int64(split[1])), decimal.NewFromInt(int64(split[2]))},
				vel: Point{decimal.NewFromInt(int64(split[3])), decimal.NewFromInt(int64(split[4])), decimal.NewFromInt(int64(split[5]))},
			})
	}

	return hails
}
