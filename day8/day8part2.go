package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point2 struct {
	Row int
	Col int
}

var inputMap2 = make(map[Point2]int)

func ScenicScore(row int, col int, maxRow int, maxCol int) int {
	point := Point2{row, col}
	height := inputMap2[point]

	// Left
	leftScore := 0
	if col == 0 {
		leftScore = 0
	} else {
		colCursor := col - 1
		for {
			farPoint := Point2{row, colCursor}
			farHeight, exists := inputMap2[farPoint]
			if !exists {
				break
			}
			leftScore += 1
			//fmt.Printf("%d, %v\n", farHeight, exists)
			if farHeight >= height {
				break
			}
			colCursor -= 1
		}
	}

	// Right
	rightScore := 0
	if col == maxCol {
		rightScore = 0
	} else {
		colCursor := col + 1
		for {
			farPoint := Point2{row, colCursor}
			farHeight, exists := inputMap2[farPoint]
			if !exists {
				break
			}
			rightScore += 1
			if farHeight >= height {
				break
			}
			colCursor += 1
		}
	}

	// Up
	upScore := 0
	if row == 0 {
		upScore = 0
	} else {
		rowCursor := row - 1
		for {
			farPoint := Point2{rowCursor, col}
			farHeight, exists := inputMap2[farPoint]
			if !exists {
				break
			}
			upScore += 1
			if farHeight >= height {
				break
			}
			rowCursor -= 1
		}
	}

	// Down
	downScore := 0
	if row == maxRow {
		downScore = 0
	} else {
		rowCursor := row + 1
		for {
			farPoint := Point2{rowCursor, col}
			farHeight, exists := inputMap2[farPoint]
			if !exists {
				break
			}
			downScore += 1
			if farHeight >= height {
				break
			}
			rowCursor += 1
		}
	}

	//fmt.Printf("(%d, %d) l: %d, r: %d, u: %d, d: %d (%d)\n", row, col, leftScore, rightScore, upScore, downScore, height)
	return leftScore * rightScore * upScore * downScore
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
		point := Point2{row, col}
		inputMap2[point] = height
		col++
	}

	// Process inputMap
	topScore := 0
	for r := 0; r < row; r++ {
		for c := 0; c < maxCol; c++ {
			score := ScenicScore(r, c, row, maxCol)
			if score > topScore {
				topScore = score
			}
		}
	}

	fmt.Println(topScore)
	// 2499 too low
}
