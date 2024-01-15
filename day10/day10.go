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
	Sum   int
}

func (cpu *CPU) Noop() {
	cpu.Clock += 1
	cpu.Check()
}

func (cpu *CPU) Addx(val int) {
	cpu.Clock += 1
	cpu.Check()
	cpu.Clock += 1
	cpu.X += val
	cpu.Check()
}

func (cpu *CPU) Check() {
	if cpu.Clock == 20 || cpu.Clock == 60 || cpu.Clock == 100 || cpu.Clock == 140 || cpu.Clock == 180 || cpu.Clock == 220 {
		cpu.Sum += cpu.Clock * cpu.X
		fmt.Printf("Cycle=%d, X=%d, Strength=%d\n", cpu.Clock, cpu.X, cpu.Clock*cpu.X)
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
	cpu := CPU{Clock: 1, X: 1}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		cpu.processLine(line)

		// 20th, 60th, 100th, 140th, 180th, and 220th
	}

	fmt.Println(cpu.Sum)
}
