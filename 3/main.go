package main

import (
	"bufio"
	"fmt"
	"strings"
)

var priorities = map[rune]int{
	'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26,
	'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52,
}

func findCommon(lines []string) rune {
	result := '0'
	var occurrence = map[rune]int{}

	for i, line := range lines {
		for _, char := range line {
			if occurrence[char] == i {
				occurrence[char] = i + 1
			}
		}
	}

	target := len(lines)
	for k, v := range occurrence {
		if v == target {
			result = k
			break
		}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(Input))
	totalPart1 := 0
	totalPart2 := 0
	lines := make([]string, 3, 3)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		com1 := findCommon([]string{line[:len(line)/2], line[len(line)/2:]})
		fmt.Printf("line: %s, common: %s\n", line, string(com1))
		totalPart1 += priorities[com1]

		lines[i%3] = line
		if (i+1)%3 == 0 {
			com2 := findCommon(lines)
			fmt.Printf("common part 2: %s\n", string(com2))
			totalPart2 += priorities[com2]
		}

	}

	fmt.Println("")
	fmt.Printf("total common: %d\n", totalPart1)
	fmt.Printf("total badges: %d\n", totalPart2)
}
