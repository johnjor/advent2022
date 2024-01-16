package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makePart1Map() map[string]int {
	values := make(map[string]int)

	// A|X = Rock, B|Y = Paper, C|Z = Scissors
	// X = +1, Y = +2, Z = +3
	// Loss = +0, Draw = +3, Win = +6
	// Left = Opponent, Right = Me
	values["A X"] = 1 + 3 // Rock, Draw
	values["A Y"] = 2 + 6 // Paper, Win
	values["A Z"] = 3 + 0 // Scissors, Loss

	values["B X"] = 1 + 0 // Rock, Loss
	values["B Y"] = 2 + 3 // Paper, Draw
	values["B Z"] = 3 + 6 // Scissors, Win

	values["C X"] = 1 + 6 // Rock, Win
	values["C Y"] = 2 + 0 // Paper, Loss
	values["C Z"] = 3 + 3 // Scissor, Draw

	return values
}

func makePart2Map() map[string]int {
	values := make(map[string]int)

	// A|X = Rock, B|Y = Paper, C|Z = Scissors
	// X = Loss, Y = Draw, Z = Win
	// Loss = +0, Draw = +3, Win = +6
	// Left = Opponent, Right = Outcome
	values["A X"] = 3 + 0 // Rock, Loss = Scissors
	values["A Y"] = 1 + 3 // Rock, Draw = Rock
	values["A Z"] = 2 + 6 // Rock, Win = Paper

	values["B X"] = 1 + 0 // Paper, Loss = Rock
	values["B Y"] = 2 + 3 // Paper, Draw = Paper
	values["B Z"] = 3 + 6 // Paper, Win = Scissors

	values["C X"] = 2 + 0 // Scissor, Loss = Paper
	values["C Y"] = 3 + 3 // Scissor, Draw = Scissors
	values["C Z"] = 1 + 6 // Scissor, Win = Rock

	return values
}

func Run() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	values := makePart2Map()

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		score += values[line]
	}

	fmt.Printf("Final score: %d\n", score)
}
