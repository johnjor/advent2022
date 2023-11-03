package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stacks := make(Stacks)
	scanner := bufio.NewScanner(file)

	InitState(stacks, scanner)

	for i := 1; i <= 9; i++ {
		v, _ := stacks[i].Peek()
		fmt.Println(v)
	}

	//for scanner.Scan() {
	//	line := strings.TrimSuffix(scanner.Text(), "\n")
	//	if !strings.HasPrefix(line, "move") {
	//		continue
	//	}
	//
	//}
}
