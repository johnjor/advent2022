package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

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

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		score += values[line]
	}

	fmt.Printf("Final score: %d\n", score)
}
