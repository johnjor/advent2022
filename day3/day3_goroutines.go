package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

func applyOffset2(ord int) int {
	if ord <= 90 {
		return ord - 38
	} else {
		return ord - 96
	}
}

func getIntersection2(line string) int {
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

func work(in <-chan string) <-chan int {
	out := make(chan int)
	go func() {
		for line := range in {
			out <- applyOffset2(getIntersection2(line))
		}
		close(out)
	}()
	return out
}

func gen(scanner *bufio.Scanner) <-chan string {
	out := make(chan string)
	go func() {
		for scanner.Scan() {
			line := strings.TrimSuffix(scanner.Text(), "\n")
			out <- line
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	in := gen(scanner)

	c1 := work(in)
	c2 := work(in)

	for n := range merge(c1, c2) {
		sum += n
	}

	fmt.Printf("%d\n", sum)
}
