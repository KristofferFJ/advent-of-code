package part1

import (
	"fmt"
	"strings"
	"testing"
)

var InputTest1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

var InputTest2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output`

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

func TestInput(t *testing.T) {
	parseInput(Input)
	low := 0
	high := 0
	var pulses []Pulse
	for i := 0; i < 1000; i++ {
		low += 1
		for _, destination := range modules["broadcaster"].destinations {
			pulses = append(pulses, Pulse{low: true, destination: destination, from: "broadcaster"})
		}
		for len(pulses) > 0 {
			pulse := pulses[0]
			pulses = pulses[1:]
			if pulse.low {
				low += 1
			} else {
				high += 1
			}

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
					pulses = append(pulses, Pulse{low: lowPulse, destination: conjunctionModule, from: pulse.destination})
				}
			}
		}
	}

	fmt.Printf("low %d, high %d = %d", low, high, low*high)
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
