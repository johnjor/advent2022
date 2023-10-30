package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TopStack struct {
	Cursor int
	Stack  [3]int
	Size   int
}

func (t *TopStack) Push(x int) {
	t.Stack[t.Cursor] = x
	t.Cursor = (t.Cursor + 1) % t.Size
}

func (t *TopStack) Current() int {
	return t.Stack[t.Cursor]
}

func (t *TopStack) Sum() int {
	sum := 0
	for _, i := range t.Stack {
		sum += i
	}
	return sum
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	n := 3
	top := TopStack{Cursor: 0, Stack: [3]int{0, 0, 0}, Size: n}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		x := scanner.Text()
		switch x {
		case "":
			if total > top.Current() {
				top.Push(total)
			}
			total = 0
		default:
			y, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			total += y
		}
	}

	if total > top.Current() {
		top.Push(total)
	}

	fmt.Printf("%d", top.Sum())
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
