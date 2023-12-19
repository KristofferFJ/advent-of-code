package part1

import (
	"fmt"
	"io.kristofferfj.github/aoc-2023-go/util"
	"strings"
	"testing"
)

var InputTest = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

type Rule struct {
	letter, returnWorkflow, operation string
	value                             int
}

var workflows = make(map[string]Workflow)

type Workflow struct {
	returnWorkflow string
	rules          []Rule
}

type Interval struct {
	minValue, maxValue int
}

func TestInput(t *testing.T) {
	parseInput(Input)
	letterInterval := make(map[string]Interval)
	letterInterval["x"] = Interval{minValue: 1, maxValue: 4000}
	letterInterval["m"] = Interval{minValue: 1, maxValue: 4000}
	letterInterval["a"] = Interval{minValue: 1, maxValue: 4000}
	letterInterval["s"] = Interval{minValue: 1, maxValue: 4000}

	fmt.Println(getValidCombinations("in", letterInterval))
}

func getValidCombinations(workflow string, letterInterval map[string]Interval) int {
	if workflow == "A" {
		sum := 1
		for _, number := range letterInterval {
			sum *= number.maxValue - number.minValue + 1
		}
		return sum
	}
	if workflow == "R" {
		return 0
	}

	sum := 0
	for _, rule := range workflows[workflow].rules {
		if rule.operation == ">" {
			validInterval := util.DuplicateMap(letterInterval)
			validInterval[rule.letter] = Interval{minValue: rule.value + 1, maxValue: letterInterval[rule.letter].maxValue}
			sum += getValidCombinations(rule.returnWorkflow, validInterval)
			letterInterval[rule.letter] = Interval{minValue: letterInterval[rule.letter].minValue, maxValue: rule.value}
		} else if rule.operation == "<" {
			validInterval := util.DuplicateMap(letterInterval)
			validInterval[rule.letter] = Interval{minValue: letterInterval[rule.letter].minValue, maxValue: rule.value - 1}
			sum += getValidCombinations(rule.returnWorkflow, validInterval)
			letterInterval[rule.letter] = Interval{minValue: rule.value, maxValue: letterInterval[rule.letter].maxValue}
		} else if rule.operation == "=" {
			validInterval := util.DuplicateMap(letterInterval)
			validInterval[rule.letter] = Interval{minValue: rule.value, maxValue: rule.value}
			sum += getValidCombinations(rule.returnWorkflow, validInterval)
			letterInterval[rule.letter] = Interval{minValue: rule.value, maxValue: rule.value}
		} else {
			panic("Invalid operation: " + rule.operation)
		}
	}

	sum += getValidCombinations(workflows[workflow].returnWorkflow, letterInterval)
	return sum
}

func parseInput(input string) {
	split := strings.Split(input, "\n\n")
	for _, workflowString := range strings.Split(split[0], "\n") {
		name := strings.Split(workflowString, "{")[0]
		rest := strings.Split(workflowString, "{")[1]
		rest = rest[:len(rest)-1]
		rulesStrings := strings.Split(rest, ",")
		mainReturnWorkflow := rulesStrings[len(rulesStrings)-1]
		rulesStrings = rulesStrings[:len(rulesStrings)-1]
		var rules []Rule
		for i := 0; i < len(rulesStrings); i++ {
			returnWorkflow := strings.Split(rulesStrings[i], ":")[1]
			ruleInfo := strings.Split(rulesStrings[i], ":")[0]
			operation := ""
			if strings.Contains(ruleInfo, ">") {
				operation = ">"
			} else if strings.Contains(ruleInfo, "<") {
				operation = "<"
			} else if strings.Contains(ruleInfo, "=") {
				operation = "="
			} else {
				panic("Invalid rule: " + ruleInfo)
			}

			rules = append(rules, Rule{
				returnWorkflow: returnWorkflow,
				letter:         ruleInfo[0:1],
				operation:      operation,
				value:          util.Int(ruleInfo[2:]),
			})
		}
		workflows[name] = Workflow{returnWorkflow: mainReturnWorkflow, rules: rules}
	}
}
