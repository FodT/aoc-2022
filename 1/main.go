package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	ElfIndex      int
	Items         []int
	TotalCalories int
}

func readFile() []elf {
	var elves []elf
	elfCount := 0

	checkFile, _ := os.Stat("./input.txt")

	if checkFile.Size() == 0 {
		return nil
	} else {
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	currentElf := elf{ElfIndex: elfCount}

	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			calories, err := strconv.Atoi(text)
			if err == nil {
				currentElf.Items = append(currentElf.Items, calories)
				currentElf.TotalCalories += calories
			}

		} else {
			elfCount++
			elves = append(elves, currentElf)
			currentElf = elf{ElfIndex: elfCount}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(currentElf.Items) > 0 {
		elves = append(elves, currentElf)
	}
	return elves
}

func topNCals(elves []elf, topN int) int {
	totalCals := 0
	for i := 0; i < topN; i++ {
		totalCals += elves[i].TotalCalories
	}

	return totalCals
}

func main() {
	elves := readFile()

	sort.Slice(elves[:], func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})

	for _, e := range elves {
		fmt.Printf("elf:\t%d \t cals: %d\n", e.ElfIndex, e.TotalCalories)
	}

	fmt.Printf("top3 total: %d\n", topNCals(elves, 3))
}
