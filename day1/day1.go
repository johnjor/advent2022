package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 1
	max_i := 1
	scanner := bufio.NewScanner(file)
	total := 0
	max_total := 0
	for scanner.Scan() {
		x := scanner.Text()
		switch x {
		case "":
			if total > max_total {
				max_total = total
				max_i = i
			}
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

	if total > max_total {
		max_total = total
		max_i = i
	}

	// Get last row?
	fmt.Printf("%d: %d\n", max_i, max_total)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
