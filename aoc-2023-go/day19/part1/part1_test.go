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
	letter, returnWorkflow string
	check                  func(int) bool
}

var workflows = make(map[string]Workflow)

type Workflow struct {
	returnWorkflow string
	rules          []Rule
}

type NumberSet struct {
	numbers []Number
}

type Number struct {
	value  int
	letter string
}

func TestInput(t *testing.T) {
	numberSets := parseInput(Input)
	sum := 0
	for _, numberSet := range numberSets {
		workflow := "in"
		for workflow != "A" && workflow != "R" {
			newWorkflowFound := false
			for _, rule := range workflows[workflow].rules {
				if newWorkflowFound {
					break
				}
				for _, number := range numberSet.numbers {
					if newWorkflowFound {
						break
					}
					if rule.letter == number.letter {
						if rule.check(number.value) {
							workflow = rule.returnWorkflow
							newWorkflowFound = true
						}
					}
				}
			}
			if !newWorkflowFound {
				workflow = workflows[workflow].returnWorkflow
			}
		}
		if workflow == "A" {
			for _, number := range numberSet.numbers {
				sum += number.value
			}
		}
	}

	fmt.Println(sum)
}

func parseInput(input string) (numberSets []NumberSet) {
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

			rules = append(rules, Rule{
				returnWorkflow: returnWorkflow,
				letter:         ruleInfo[0:1],
				check:          getFunc(ruleInfo[1:]),
			})
		}
		workflows[name] = Workflow{returnWorkflow: mainReturnWorkflow, rules: rules}
	}

	for _, numberString := range strings.Split(split[1], "\n") {
		var numbers []Number
		numberString = numberString[1 : len(numberString)-1]
		for _, individualNumber := range strings.Split(numberString, ",") {
			numbers = append(numbers, Number{
				value:  util.Int(strings.Split(individualNumber, "=")[1]),
				letter: strings.Split(individualNumber, "=")[0],
			})
		}
		numberSets = append(numberSets, NumberSet{numbers: numbers})
	}

	return numberSets
}

func getFunc(rule string) func(int) bool {
	if strings.Contains(rule, ">") {
		return func(value int) bool {
			return value > util.Int(rule[1:])
		}
	} else if strings.Contains(rule, "<") {
		return func(value int) bool {
			return value < util.Int(rule[1:])
		}
	} else if strings.Contains(rule, "=") {
		return func(value int) bool {
			return value == util.Int(rule[1:])
		}
	}
	panic("Invalid rule: " + rule)
}
