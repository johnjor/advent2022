package day5

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stacks map[int]*Stack

type Stack struct {
	top    *node
	length int
}

type node struct {
	value string
	prev  *node
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(value string) *Stack {
	n := &node{value, s.top}
	s.top = n
	s.length++
	return s
}

func (s *Stack) Pop() (string, error) {
	if s.length == 0 {
		return "", errors.New("empty stack")
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value, nil
}

func (s *Stack) Peek() (string, error) {
	if s.length == 0 {
		return "", errors.New("empty stack")
	}

	return s.top.value, nil
}

type Move struct {
	Qty  int
	From int
	To   int
}

func InitState(stacks Stacks, scanner *bufio.Scanner) {
	for i := 1; i <= 9; i++ {
		stacks[i] = NewStack()
	}

	//lines := make([]string, 8)
	//
	//for i := 0; i < 7; i++ {
	//	lines[i] = scanner.Text()
	//}
	//
	//column := 2
	//
	//for i := 0; i < 9; i++ {
	//	for j := 7; j >= 0; j-- {
	//		c := string(lines[j][column])
	//		if c != " " {
	//			stacks[i].Push(string(c))
	//		}
	//	}
	//}

	stacks[1].Push("N").Push("S").Push("D").Push("C").Push("V").Push("Q").Push("T")
	stacks[2].Push("M").Push("F").Push("V")
	stacks[3].Push("F").Push("Q").Push("W").Push("D").Push("P").Push("N").Push("H").Push("M")
	stacks[4].Push("D").Push("Q").Push("R").Push("T").Push("F")
	stacks[5].Push("R").Push("F").Push("M").Push("N").Push("Q").Push("H").Push("V").Push("B")
	stacks[6].Push("C").Push("F").Push("G").Push("N").Push("P").Push("W").Push("Q")
	stacks[7].Push("W").Push("F").Push("R").Push("L").Push("C").Push("T")
	stacks[8].Push("T").Push("Z").Push("N").Push("S")
	stacks[9].Push("M").Push("S").Push("D").Push("J").Push("R").Push("Q").Push("H").Push("N")
}

func ParseMove(line string) *Move {
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := r.FindStringSubmatch(line)
	if matches == nil || len(matches) != 4 {
		panic(fmt.Sprintf("line: %v, matches %v", line, matches))
	}

	q, _ := strconv.Atoi(matches[1])
	f, _ := strconv.Atoi(matches[2])
	t, _ := strconv.Atoi(matches[3])

	return &Move{q, f, t}
}

func Run() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stacks := make(Stacks)
	scanner := bufio.NewScanner(file)

	InitState(stacks, scanner)

	//for i := 1; i <= 9; i++ {
	//	v, _ := stacks[i].Peek()
	//	fmt.Println(v)
	//}

	crane := NewStack()

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		if !strings.HasPrefix(line, "move") {
			continue
		}

		move := ParseMove(line)

		for i := 1; i <= move.Qty; i++ {
			v, err := stacks[move.From].Pop()
			if err != nil {
				panic(err)
			}
			crane.Push(v)
		}

		for i := 1; i <= move.Qty; i++ {
			v, err := crane.Pop()
			if err != nil {
				panic(err)
			}
			stacks[move.To].Push(v)
		}
	}

	buf := strings.Builder{}
	for i := 1; i <= 9; i++ {
		v, err := stacks[i].Peek()
		if err != nil {
			panic(err)
		}
		buf.WriteString(v)
	}

	fmt.Println(buf.String())
}
