package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	Clock int
	X     int
}

func (cpu *CPU) Noop() {
	cpu.Tick()
}

func (cpu *CPU) Addx(val int) {
	cpu.Tick()
	cpu.Tick()
	cpu.X += val
}

func (cpu *CPU) Tick() {
	if cpu.Clock%40 == cpu.X-1 || cpu.Clock%40 == cpu.X || cpu.Clock%40 == cpu.X+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	cpu.Clock += 1
	if cpu.Clock%40 == 0 {
		fmt.Print("\n")
	}
}

func (cpu *CPU) processLine(line string) {
	if line == "noop" {
		cpu.Noop()
		return
	}

	split := strings.Split(line, " ")
	if split[0] == "addx" {
		x, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		cpu.Addx(x)
	}
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cpu := CPU{Clock: 0, X: 1}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		cpu.processLine(line)
	}
}
