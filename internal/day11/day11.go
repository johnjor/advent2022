package day11

import (
	"fmt"
	"os"
	"regexp"
	"sort"
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
	IfFalse   *Monkey
	Count     int
}

func (monkey *Monkey) NextItem() int {
	item := monkey.Items[0]
	monkey.Items = monkey.Items[1:]
	return item
}

func (monkey *Monkey) HandleNextItem() {
	item := monkey.NextItem()
	monkey.Count++
	item = monkey.Operation(item) // Apply op
	item = item / 3               // Relief
	if item%monkey.Test == 0 {    // Test
		// Throw to IfTrue
		monkey.IfTrue.Items = append(monkey.IfTrue.Items, item)
	} else {
		// Throw to IfFalse
		monkey.IfFalse.Items = append(monkey.IfFalse.Items, item)
	}
}

func (monkey *Monkey) Turn() {
	for len(monkey.Items) > 0 {
		monkey.HandleNextItem()
	}
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

	// Parse data into structs
	for _, m := range monkeyData {
		monkeys = append(monkeys, MakeMonkey(m))
	}

	// Map int ids to references
	for i, _ := range monkeys {
		ifTrueId := monkeys[i].IfTrueId
		ifFalseId := monkeys[i].IfFalseId
		monkeys[i].IfTrue = monkeys[ifTrueId]
		monkeys[i].IfFalse = monkeys[ifFalseId]
	}

	//fmt.Println("Initial")
	//for _, m := range monkeys {
	//	fmt.Printf("%+v\n", m)
	//	//fmt.Printf("%d true:%d, false:%d\n", m.Id, m.IfTrue.Id, m.IfFalse.Id)
	//}

	for i := 1; i <= 20; i++ {
		for _, m := range monkeys {
			m.Turn()
		}
		//fmt.Printf("After %d round\n", i)
		//for _, m := range monkeys {
		//	fmt.Printf("%+v\n", m)
		//}
	}

	//for _, m := range monkeys {
	//	fmt.Printf("Id=%d, Count=%d\n", m.Id, m.Count)
	//}

	sort.Slice(monkeys[:], func(i, j int) bool {
		return monkeys[i].Count > monkeys[j].Count
	})

	fmt.Printf("%d", monkeys[0].Count*monkeys[1].Count)
}
