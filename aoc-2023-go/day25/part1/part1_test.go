package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"math"
	"sort"
	"strings"
	"testing"
)

var InputTest = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

type Edge struct{ one, two string }

var connections = make(map[string][]string)
var vertices []string
var edges []Edge

type EdgeCount struct {
	edge  Edge
	count int
}

var edgeCounts []EdgeCount
var paths [][]string
var djikstra = make(map[string]int)

func TestInput(t *testing.T) {
	for _, line := range strings.Split(Input, "\n") {
		parts := strings.Split(line, ": ")
		if !util.Contains(vertices, parts[0]) {
			vertices = append(vertices, parts[0])
			djikstra[parts[0]] = math.MaxInt
		}
		for _, conn := range strings.Split(parts[1], " ") {
			connections[parts[0]] = append(connections[parts[0]], conn)
			connections[conn] = append(connections[conn], parts[0])
			edges = append(edges, Edge{one: parts[0], two: conn})
			if !util.Contains(vertices, conn) {
				vertices = append(vertices, conn)
				djikstra[conn] = math.MaxInt
			}
		}
	}

	for i := 0; i < len(vertices)-1; i++ {
		paths = append(paths, getPaths(i)...)
	}

	for _, edge := range edges {
		count := 0
		for _, path := range paths {
			if util.Contains(path, edge.one) && util.Contains(path, edge.two) {
				count++
			}
		}
		edgeCounts = append(edgeCounts, EdgeCount{edge: edge, count: count})
	}

	sort.Slice(edgeCounts, func(i, j int) bool {
		return edgeCounts[i].count > edgeCounts[j].count
	})

	for i := 0; i < len(edgeCounts); i++ {
		for j := i + 1; j < len(edgeCounts); j++ {
			for k := j + 1; k < len(edgeCounts); k++ {
				duplicate := util.DuplicateMapListValues(connections)
				duplicate = removeConnection(duplicate, edgeCounts[i].edge)
				duplicate = removeConnection(duplicate, edgeCounts[j].edge)
				duplicate = removeConnection(duplicate, edgeCounts[k].edge)
				groups := getGroups(duplicate)
				if len(groups) > 1 {
					fmt.Println(len(groups[0]) * len(groups[1]))
					return
				}
			}
		}
	}

	fmt.Println(paths)
}

func getPaths(index int) (paths [][]string) {
	thisDjikstra := util.DuplicateMap(djikstra)
	vertex := vertices[index]
	thisDjikstra[vertex] = 0
	next := []string{vertex}
	for len(next) > 0 {
		current := next[0]
		next = next[1:]
		for _, conn := range connections[current] {
			if thisDjikstra[conn] > thisDjikstra[current]+1 {
				thisDjikstra[conn] = thisDjikstra[current] + 1
				next = append(next, conn)
			}
		}
	}

	for j := index + 1; j < len(vertices); j++ {
		toVertex := vertices[j]
		path := []string{toVertex}
		dist := thisDjikstra[path[len(path)-1]]
		for dist > 0 {
			for _, conn := range connections[path[len(path)-1]] {
				if thisDjikstra[conn] == dist-1 {
					path = append(path, conn)
					dist--
					break
				}
			}
		}
		paths = append(paths, path)
	}

	return paths
}

func removeConnection(conns map[string][]string, edge Edge) map[string][]string {
	conns[edge.one] = util.RemoveElement(conns[edge.one], edge.two)
	conns[edge.two] = util.RemoveElement(conns[edge.two], edge.one)
	return conns
}

func getGroups(connections map[string][]string) [][]string {
	var groups [][]string
	inGroup := make(map[string]bool)
	for key, value := range connections {
		if inGroup[key] {
			continue
		}
		group := []string{key}
		inGroup[key] = true
		newConns := value
		for len(newConns) > 0 {
			conn := newConns[0]
			newConns = newConns[1:]
			if inGroup[conn] {
				continue
			}
			group = append(group, conn)
			inGroup[conn] = true
			newConns = append(newConns, connections[conn]...)
		}
		groups = append(groups, group)
	}
	return groups
}
