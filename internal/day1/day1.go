package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 1
	max_i := 1
	_ = max_i
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
	//fmt.Printf("%d: %d\n", max_i, max_total)
	// Question specifically asks for just the Calories as output
	fmt.Printf("%d\n", max_total)
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
