package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	in := strings.Split(Input, "\n\n")
	board := in[0]
	moves := in[1]

	boardlines := strings.Split(board, "\n")
	var stacks []Stack
	var stacks2 []Stack
	for i := len(boardlines) - 1; i >= 0; i-- {
		line := boardlines[i]
		if i == len(boardlines)-1 {
			line = strings.TrimSpace(boardlines[i])
			numStacks, _ := strconv.Atoi(string(line[len(line)-1]))
			fmt.Printf("creating %d stacks\n", numStacks)
			stacks = make([]Stack, numStacks, numStacks)
			stacks2 = make([]Stack, numStacks, numStacks)
		} else {
			k := 0
			for j := 1; j < len(line); j += 4 {

				if !unicode.IsSpace(rune(line[j])) {
					stacks[k].Push(string(line[j]))
					stacks2[k].Push(string(line[j]))
				}
				k++
			}
		}

	}

	fmt.Println(stacks)

	scanner := bufio.NewScanner(strings.NewReader(moves))
	for scanner.Scan() {
		line := scanner.Text()
		var numMoves, from, to int
		//move 2 from 4 to 6
		_, _ = fmt.Fscanf(strings.NewReader(line), "move %d from %d to %d\n", &numMoves, &from, &to)
		for i := 0; i < numMoves; i++ {
			val, _ := stacks[from-1].Pop()
			stacks[to-1].Push(val)
		}

		val, _ := stacks2[from-1].PopN(numMoves)
		stacks2[to-1].PushN(val)

	}

	fmt.Print("stack 1 state: ")
	for i := 0; i < len(stacks); i++ {
		val, had := stacks[i].Pop()
		if had {
			fmt.Print(val)
			stacks[i].Push(val)
		}
	}
	fmt.Println("")

	fmt.Print("stack 2 state: ")
	for i := 0; i < len(stacks2); i++ {
		val, had := stacks2[i].Pop()
		if had {
			fmt.Print(val)
			stacks2[i].Push(val)
		}
	}
	fmt.Println("")

}
