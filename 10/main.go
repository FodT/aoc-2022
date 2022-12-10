package main

import (
	"bufio"
	"fmt"
	"strings"
)

var programCounter int
var regX int

func main() {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	total := 0
	regX = 1
	for scanner.Scan() {

		line := scanner.Text()
		var val int
		ticks := 1
		if !strings.HasPrefix(line, "n") {
			ticks = 2
			fmt.Sscanf(line, "addx %d", &val)
		}
		total += addX(ticks, val)
	}

	fmt.Printf("\n\nPart 1 Total: %d\n", total)
}

func addX(ticks int, incr int) int {
	signal := 0
	for i := 0; i < ticks; i++ {
		position := programCounter % 40
		if position == 0 {
			fmt.Println("")
		}
		programCounter++
		// draw pixel
		if Abs(position-regX) < 2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if (programCounter-20)%40 == 0 {
			signal = regX * programCounter
		}
	}
	regX += incr
	return signal
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
