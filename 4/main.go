package main

import (
	"bufio"
	"fmt"
	"strings"
)

func contained(aMin int, aMax int, bMin int, bMax int) bool {
	if aMin <= bMin {
		if aMax >= bMax {
			return true
		}
	}

	if bMin <= aMin {
		if bMax >= aMax {
			return true
		}
	}
	return false
}

func overlaps(aMin int, aMax int, bMin int, bMax int) bool {
	if aMin <= bMax && bMin <= aMax {
		return true
	}
	return false
}

func main() {
	totalContained := 0
	totalOverlaps := 0
	scanner := bufio.NewScanner(strings.NewReader(Input))
	for scanner.Scan() {
		var aMin, aMax, bMin, bMax int
		fmt.Fscanf(strings.NewReader(scanner.Text()), "%d-%d,%d-%d\n", &aMin, &aMax, &bMin, &bMax)

		if contained(aMin, aMax, bMin, bMax) {
			totalContained++
		}
		if overlaps(aMin, aMax, bMin, bMax) {
			totalOverlaps++
		}
	}

	fmt.Printf("totalContained: %d\n", totalContained)
	fmt.Printf("totalOverlaps: %d\n", totalOverlaps)
}
