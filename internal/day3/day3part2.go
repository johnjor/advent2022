package day3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	for _, r := range []rune(s) {
		set[r] = true
	}
	return set
}

func setIntersection(set_a map[rune]bool, set_b map[rune]bool) map[rune]bool {
	intersection := make(map[rune]bool)
	for r, _ := range set_a {
		if set_b[r] {
			intersection[r] = true
		}
	}
	return intersection
}

func applyOffset3(ord int) int {
	if ord <= 90 {
		return ord - 38
	} else {
		return ord - 96
	}
}

func getCommonBadgeValue(bag1 string, bag2 string, bag3 string) int {
	set_a := makeSet(bag1)
	set_b := makeSet(bag2)
	set_c := makeSet(bag3)

	d := setIntersection(setIntersection(set_a, set_b), set_c)

	for k := range d {
		return applyOffset3(int(k))
	}
	return -1
}

func RunPart2() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	bags := [3]string{}
	cursor := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bags[cursor] = strings.TrimSuffix(scanner.Text(), "\n")
		cursor = (cursor + 1) % 3

		if cursor == 0 {
			sum += getCommonBadgeValue(bags[0], bags[1], bags[2])
		}
	}

	fmt.Printf("%d\n", sum)
}
