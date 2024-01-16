package day3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode/utf8"
)

func applyOffset(ord int) int {
	if ord <= 90 {
		return ord - 38
	} else {
		return ord - 96
	}
}

func getIntersection(line string) int {
	length := utf8.RuneCountInString(line)
	half := length / 2
	cursor1 := 0
	cursor2 := 0

	bag1 := []rune(line[:half])
	bag2 := []rune(line[half:])

	sort.Slice(bag1, func(i int, j int) bool { return bag1[i] < bag1[j] })
	sort.Slice(bag2, func(i int, j int) bool { return bag2[i] < bag2[j] })

	for cursor1 < half && cursor2 < half {
		a := int(bag1[cursor1])
		b := int(bag2[cursor2])
		if a == b {
			return a
		} else if a < b {
			cursor1++
		} else {
			cursor2++
		}
	}
	fmt.Printf("%v, %v\n", bag1, bag2)
	return -1
}

func Run() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")

		sum += applyOffset(getIntersection(line))
	}

	fmt.Printf("%d\n", sum)
}
