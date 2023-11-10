package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	Row int
	Col int
}

var inputMap = make(map[Point]int)
var maskMap = make(map[Point]int)

func runRight(maxRow int, maxCol int) {
	maxHeight := -1
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			point := Point{r, c}
			currentHeight := inputMap[point]
			if currentHeight > maxHeight {
				maskMap[point] = 1
				maxHeight = currentHeight
			}
		}
		maxHeight = -1
	}
}

func runLeft(maxRow int, maxCol int) {
	maxHeight := -1
	for r := 0; r < maxRow; r++ {
		for c := maxCol - 1; c >= 0; c-- {
			point := Point{r, c}
			currentHeight := inputMap[point]
			if currentHeight > maxHeight {
				maskMap[point] = 1
				maxHeight = currentHeight
			}
		}
		maxHeight = -1
	}
}

func runDown(maxRow int, maxCol int) {
	maxHeight := -1
	for c := 0; c < maxCol; c++ {
		for r := 0; r < maxRow; r++ {
			point := Point{r, c}
			currentHeight := inputMap[point]
			if currentHeight > maxHeight {
				maskMap[point] = 1
				maxHeight = currentHeight
			}
		}
		maxHeight = -1
	}
}

func runUp(maxRow int, maxCol int) {
	maxHeight := -1
	for c := 0; c < maxCol; c++ {
		for r := maxRow - 1; r >= 0; r-- {
			point := Point{r, c}
			currentHeight := inputMap[point]
			if currentHeight > maxHeight {
				maskMap[point] = 1
				maxHeight = currentHeight
			}
		}
		maxHeight = -1
	}
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Populate inputMap
	reader := bufio.NewReader(file)
	row, col, maxCol := 0, 0, 0
	for {
		r, _, err := reader.ReadRune()
		if r == '\n' {
			row++
			maxCol = col
			col = 0
			continue
		}
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		height := int(r - '0')
		point := Point{row, col}
		inputMap[point] = height
		col++
	}

	runRight(row, maxCol)
	runLeft(row, maxCol)
	runUp(row, maxCol)
	runDown(row, maxCol)

	total := 0
	for _, v := range maskMap {
		total += v
	}

	fmt.Println(total)
}
