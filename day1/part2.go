package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func lowest(arr []int) (int, int) {
	lowestI := 0
	lowestV := math.MaxInt
	for i, v := range arr {
		if v < lowestV {
			lowestI = i
			lowestV = v
		}
	}
	return lowestI, lowestV
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	top := []int{0, 0, 0}

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		x := scanner.Text()
		switch x {
		case "":
			var lowest_i, lowest_v = lowest(top)
			if total > lowest_v {
				top[lowest_i] = total
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

	var lowestI, lowestV = lowest(top)
	if total > lowestV {
		top[lowestI] = lowestV
	}

	fmt.Printf("%d\n", sum(top))
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
