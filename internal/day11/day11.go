package day11

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Op func(int) int

type Monkey struct {
	Id        int
	Items     []int
	Operation Op
	Test      int
	IfTrueId  int
	IfFalseId int
	IfTrue    *Monkey
	ifFalse   *Monkey
}

func MakeOp(str string) Op {
	if str == "old * old" {
		//fmt.Printf("%s -> square\n", str)
		return func(old int) int {
			return old * old
		}
	}

	pattern := regexp.MustCompile(`old (.) (\d+)`)
	matches := pattern.FindStringSubmatch(str)
	operand, _ := strconv.Atoi(matches[2])
	if matches[1] == "+" {
		//fmt.Printf("%s -> + %d\n", str, operand)
		return func(old int) int {
			return old + operand
		}
	} else if matches[1] == "*" {
		//fmt.Printf("%s -> * %d\n", str, operand)
		return func(old int) int {
			return old * operand
		}
	} else {
		panic(fmt.Sprintf("Not + or *: %s", str))
	}
}

func MakeMonkey(str string) *Monkey {
	idPattern := regexp.MustCompile(`Monkey (\d+):`)
	idMatches := idPattern.FindStringSubmatch(str)
	monkeyId, _ := strconv.Atoi(idMatches[1])

	itemsPattern := regexp.MustCompile(`Starting items: (.+)\n`)
	itemsMatches := itemsPattern.FindStringSubmatch(str)
	monkeyItems := make([]int, 0)
	for _, i := range strings.Split(itemsMatches[1], ", ") {
		x, _ := strconv.Atoi(i)
		monkeyItems = append(monkeyItems, x)
	}

	opPattern := regexp.MustCompile(`Operation: new = (.+)\n`)
	opMatches := opPattern.FindStringSubmatch(str)
	monkeyOp := MakeOp(opMatches[1])

	testPattern := regexp.MustCompile(`divisible by (\d+)`)
	testMatches := testPattern.FindStringSubmatch(str)
	monkeyTest, _ := strconv.Atoi(testMatches[1])

	ifTruePattern := regexp.MustCompile(`If true: throw to monkey (\d+)`)
	ifTrueMatches := ifTruePattern.FindStringSubmatch(str)
	ifTrue, _ := strconv.Atoi(ifTrueMatches[1])

	ifFalsePattern := regexp.MustCompile(`If false: throw to monkey (\d+)`)
	ifFalseMatches := ifFalsePattern.FindStringSubmatch(str)
	ifFalse, _ := strconv.Atoi(ifFalseMatches[1])

	return &Monkey{
		Id:        monkeyId,
		Items:     monkeyItems,
		Operation: monkeyOp,
		Test:      monkeyTest,
		IfTrueId:  ifTrue,
		IfFalseId: ifFalse,
	}
}

func Run(filename string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	contents := string(b)

	monkeyData := strings.Split(contents, "\n\n")
	monkeys := make([]*Monkey, 0)

	for _, m := range monkeyData {
		monkeys = append(monkeys, MakeMonkey(m))
	}

	for _, m := range monkeys {
		fmt.Printf("%+v\n", m)
	}

}
