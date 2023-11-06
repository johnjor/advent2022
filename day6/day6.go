package main

import (
	"bufio"
	"fmt"
	"os"
)

type Window struct {
	buffer []rune
	sums   map[rune]int
	length int
}

func NewWindow(n int) *Window {
	b := make([]rune, n)
	s := make(map[rune]int)
	return &Window{b, s, n}
}

func (w *Window) Push(in rune) {
	out := w.buffer[0]
	for i := 1; i < w.length; i++ {
		w.buffer[i-1] = w.buffer[i]
	}
	w.buffer[w.length-1] = in

	w.sums[out] -= 1
	_, present := w.sums[in]
	if !present {
		w.sums[in] = 1
	} else {
		w.sums[in] += 1
	}
}

func (w *Window) AreAllUnique() bool {
	for _, count := range w.sums {
		if count > 1 {
			return false
		}
	}
	return true
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	counter := 0
	n := 14
	window := NewWindow(n)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}
		counter++

		window.Push(r)

		if counter < 4 {
			continue
		}

		if !window.AreAllUnique() {
			continue
		} else {
			break
		}
	}

	fmt.Println(counter)
}
