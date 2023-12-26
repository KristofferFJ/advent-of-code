package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"sort"
	"strings"
	"testing"
)

var InputTest = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

type Brick struct{ points []*Point }
type Point struct{ x, y, z int }
type Area struct{ x, y int }

func TestInput(t *testing.T) {
	bricks := parseInput(Input)
	fall(bricks)
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].points[0].z < bricks[j].points[0].z
	})
	supporting := map[int][]int{}
	supportedBy := map[int][]int{}
	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		pointsMap := make(map[Point]bool)
		for _, point := range brick.points {
			pointsMap[Point{x: point.x, y: point.y, z: point.z}] = true
		}
		for j := i + 1; j < len(bricks); j++ {
			otherBrick := bricks[j]
			if otherBrick.points[0].z == 1 {
				continue
			}
			for _, point := range otherBrick.points {
				if pointsMap[Point{x: point.x, y: point.y, z: point.z - 1}] {
					supporting[i] = append(supporting[i], j)
					supportedBy[j] = append(supportedBy[j], i)
					break
				}
			}
		}
	}

	removable := 0
	for i := 0; i < len(bricks); i++ {
		if len(supporting[i]) == 0 {
			removable++
			continue
		}
		moreSupporters := true
		for _, j := range supporting[i] {
			if len(supportedBy[j]) == 1 {
				moreSupporters = false
				break
			}
		}
		if moreSupporters {
			removable++
		}
	}

	fmt.Println(removable)
}

func fall(bricks []*Brick) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].points[0].z < bricks[j].points[0].z
	})
	for i := 0; i < len(bricks); i++ {
		brick := *bricks[i]
		alt := 1
		area := make(map[Area]bool)
		for _, point := range brick.points {
			area[Area{x: point.x, y: point.y}] = true
		}
		for j := 0; j < i; j++ {
			for _, point := range bricks[j].points {
				if area[Area{x: point.x, y: point.y}] {
					alt = util.Max(point.z+1, alt)
				}
			}
		}
		diff := brick.points[0].z - alt
		for _, point := range brick.points {
			point.z -= diff
		}
	}
}

func parseInput(input string) (bricks []*Brick) {
	for _, line := range strings.Split(input, "\n") {
		nums := util.IntArray(line)
		brick := Brick{}
		for x := nums[0]; x <= nums[3]; x++ {
			for y := nums[1]; y <= nums[4]; y++ {
				for z := nums[2]; z <= nums[5]; z++ {
					brick.points = append(brick.points, &Point{x: x, y: y, z: z})
				}
			}
		}
		bricks = append(bricks, &brick)
	}

	return bricks
}
