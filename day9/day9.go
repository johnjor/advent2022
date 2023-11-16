package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func (origin Coordinate) CalculateOffset(other Coordinate) Coordinate {
	xDistance := other.X - origin.X
	yDistance := other.Y - origin.Y
	return Coordinate{xDistance, yDistance}
}

type Board struct {
	History map[Coordinate]int
	Head    Coordinate
	Tail    Coordinate
}

func NewBoard() *Board {
	history := make(map[Coordinate]int)
	zero := Coordinate{0, 0}
	history[zero] = 1 // record initial location of tail in history
	return &Board{history, zero, zero}
}

func (board *Board) Move(direction string) {
	switch direction {
	case "L":
		board.Head.X--
	case "R":
		board.Head.X++
	case "U":
		board.Head.Y++
	case "D":
		board.Head.Y--
	}
	board.moveTail()
}

func (board *Board) moveTail() {
	// move tail to follow head
	// tail x,y - head x,y gives an offset, then subtract that offset against tail's position to move it
	offset := board.Head.CalculateOffset(board.Tail)
	//fmt.Printf("Head: %v, Tail: %v, Offset: %v\n", board.Head, board.Tail, offset)
	if Abs(offset.X) > 1 || Abs(offset.Y) > 1 {
		board.Tail.X -= ZeroOrOne(offset.X)
		board.Tail.Y -= ZeroOrOne(offset.Y)

		// record current location of tail in history
		board.History[board.Tail] = 1
		//fmt.Printf("Moved Tail to %v\n", board.Tail)
	}

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ZeroOrOne(x int) int {
	if x == 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	if x < 1 {
		return -1
	}
	return x
}

func parseLine(line string) (string, int) {
	r := regexp.MustCompile(`([UDLR]) (\d+)`)
	matches := r.FindStringSubmatch(line)

	direction := matches[1]
	distance, _ := strconv.Atoi(matches[2])
	return direction, distance
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	board := NewBoard()

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		direction, distance := parseLine(line)
		//fmt.Printf("Direction: %s, Distance: %d\n", direction, distance)
		for i := 0; i < distance; i++ {
			board.Move(direction)
		}
	}

	// Count spaces where tail has visited
	sum := 0
	for _, v := range board.History {
		sum += v
	}
	fmt.Println(sum)
}
