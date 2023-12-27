package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
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

type Conn struct {
	one, two string
}

var connList []Conn

func TestInput(t *testing.T) {
	connections := make(map[string][]string)
	for _, line := range strings.Split(Input, "\n") {
		parts := strings.Split(line, ": ")
		for _, conn := range strings.Split(parts[1], " ") {
			connections[parts[0]] = append(connections[parts[0]], conn)
			connections[conn] = append(connections[conn], parts[0])
		}
	}
	done := make(map[string]bool)
	for key, value := range connections {
		done[key] = true
		for _, conn := range value {
			if done[conn] {
				continue
			}
			connList = append(connList, Conn{one: key, two: conn})
		}
	}

	for i := 0; i < len(connList); i++ {
		for j := i + 1; j < len(connList); j++ {
			for k := j + 1; k < len(connList); k++ {
				duplicate := util.DuplicateMapListValues(connections)
				duplicate = removeConnection(duplicate, connList[i])
				duplicate = removeConnection(duplicate, connList[j])
				duplicate = removeConnection(duplicate, connList[k])
				groups := getGroups(duplicate)
				if len(groups) > 1 {
					fmt.Println(len(groups[0]) * len(groups[1]))
				}
			}
		}
	}
}

func countConnections(connections map[string][]string) {
	count := 0
	for _, value := range connections {
		count += len(value)
	}
	fmt.Printf("%d connections\n", count/2)
}

func removeConnection(conns map[string][]string, conn Conn) map[string][]string {
	conns[conn.one] = util.RemoveElement(conns[conn.one], conn.two)
	conns[conn.two] = util.RemoveElement(conns[conn.two], conn.one)
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
