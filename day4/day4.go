package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compareValues(line string) int {
	split := strings.Split(line, ",")
	elf1 := strings.Split(split[0], "-")
	elf2 := strings.Split(split[1], "-")

	a, err := strconv.Atoi(elf1[0])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(elf1[1])
	if err != nil {
		panic(err)
	}

	x, err := strconv.Atoi(elf2[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(elf2[1])
	if err != nil {
		panic(err)
	}

	if (a >= x && b <= y) || (a <= x && b >= y) {
		//fmt.Printf("YES: %v %v\n", elf1, elf2)
		return 1
	} else {
		//fmt.Printf("NO: %v %v\n", elf1, elf2)
		return 0
	}
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		sum += compareValues(line)
	}

	fmt.Println(sum)
}
