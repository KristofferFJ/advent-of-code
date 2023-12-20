package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

type ModuleInfo struct {
	moduleType   string
	destinations []string
	on           bool
}

var conjunctions = make(map[string][]*ModulePulse)
var modules = make(map[string]*ModuleInfo)

type ModulePulse struct {
	name string
	low  bool
}

type Pulse struct {
	low         bool
	from        string
	destination string
}

var loop = make(map[string][]int)

func TestInput(t *testing.T) {
	parseInput(Input)
	var pulses []Pulse
	foundAt := 0
	for i := 0; foundAt == 0; i++ {
		for _, destination := range modules["broadcaster"].destinations {
			pulses = append(pulses, Pulse{low: true, destination: destination, from: "broadcaster"})
		}
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]

			module, ok := modules[pulse.destination]
			if ok && module.moduleType == "%" {
				if !pulse.low {
					continue
				}
				module.on = !module.on
				for _, destination := range module.destinations {
					pulses = append(pulses, Pulse{low: !module.on, destination: destination, from: pulse.destination})
				}
			} else if ok && module.moduleType == "&" {
				if pulse.destination == "bb" && !pulse.low {
					loop[pulse.from] = append(loop[pulse.from], i)
					loopFound, lcm := loopDetected()
					if loopFound {
						foundAt = i
						fmt.Printf("loop found at %d, lcm is %d\n", i, lcm)
					}
				}
				for _, modulePulse := range conjunctions[pulse.destination] {
					if modulePulse.name == pulse.from {
						modulePulse.low = pulse.low
					}
				}
				lowPulse := true
				for _, modulePulse := range conjunctions[pulse.destination] {
					if modulePulse.low {
						lowPulse = false
					}
				}
				for _, conjunctionModule := range module.destinations {
					pulses = append(pulses, Pulse{
						low:         lowPulse,
						destination: conjunctionModule,
						from:        pulse.destination,
					})
				}
			}
		}
	}
}

func loopDetected() (bool, int) {
	for _, ints := range loop {
		if len(ints) < 3 {
			return false, 0
		}
	}
	var keys []string
	for key, _ := range loop {
		keys = append(keys, key)
	}
	var values []int
	for _, value := range loop {
		top := len(value) - 1
		values = append(values, value[top]-value[top-1])
		if value[top]-value[top-1] != value[top-1]-value[top-2] {
			return false, 0
		}
	}

	return true, util.LCMArray(values)
}

func parseInput(input string) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		moduleType := "none"
		name := parts[0]
		if parts[0][0:1] == "%" {
			moduleType = parts[0][0:1]
			name = parts[0][1:]
		} else if parts[0][0:1] == "&" {
			moduleType = parts[0][0:1]
			name = parts[0][1:]
			conjunctions[name] = []*ModulePulse{}
		}
		var destinationModules []string
		for _, part := range parts[2:] {
			trimmed := strings.Trim(part, ",")
			destinationModules = append(destinationModules, trimmed)
		}
		modules[name] = &ModuleInfo{moduleType: moduleType, destinations: destinationModules}
	}

	for conjunctionName, _ := range conjunctions {
		for moduleName, module := range modules {
			for _, destinationModule := range module.destinations {
				if destinationModule == conjunctionName {
					conjunctions[conjunctionName] = append(conjunctions[conjunctionName],
						&ModulePulse{name: moduleName, low: true},
					)
				}
			}
		}
	}
}
