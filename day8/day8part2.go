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

	colCursor := col - 1
	leftScore := 1
	for {
		farPoint := Point2{row, colCursor}
		farHeight, exists := inputMap2[farPoint]
		if !exists || farHeight >= height {
			break
		}
		leftScore += 1
		colCursor -= 1
	}

	colCursor = col + 1
	rightScore := 1
	for {
		farPoint := Point2{row, colCursor}
		farHeight, exists := inputMap2[farPoint]
		if !exists || farHeight >= height {
			break
		}
		rightScore += 1
		colCursor += 1
	}

	rowCursor := row - 1
	upScore := 1
	for {
		farPoint := Point2{row, colCursor}
		farHeight, exists := inputMap2[farPoint]
		if !exists || farHeight >= height {
			break
		}
		upScore += 1
		rowCursor -= 1
	}

	rowCursor = row + 1
	downScore := 1
	for {
		farPoint := Point2{row, colCursor}
		farHeight, exists := inputMap2[farPoint]
		if !exists || farHeight >= height {
			break
		}
		downScore += 1
		rowCursor += 1
	}

	//fmt.Printf("l: %d, r: %d, u: %d, d: %d\n", leftScore, rightScore, upScore, downScore)
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
