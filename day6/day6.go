package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	counter := 0
	var slot1 rune = 0
	var slot2 rune = 0
	var slot3 rune = 0
	var slot4 rune = 0

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		counter++

		slot1, slot2, slot3, slot4 = slot2, slot3, slot4, r

		if counter < 4 {
			continue
		}

		if slot1 == slot2 || slot1 == slot3 || slot1 == slot4 || slot2 == slot3 || slot2 == slot4 || slot3 == slot4 {
			continue
		} else {
			break
		}
	}

	fmt.Println(counter)
}
