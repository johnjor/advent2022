package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 1

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		x := scanner.Text()
		switch x {
		case "":
			fmt.Printf("%d: %d\n", i, total)
			i++
			total = 0
		default:
			y, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			total += y
		}
	}

	// Get last row?
	fmt.Printf("%d: %d\n", i, total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
