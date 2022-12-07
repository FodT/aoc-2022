package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile() (int, int) {
	// Rock A/X: 1
	// Paper B/Y: 2
	// Scissors C/Z: 3
	// Loss: 0
	// Draw: 3
	// Win: 6
	scores := map[string]int{
		"A X": 1 + 3, // rock.rock -  draw, 1+3
		"A Y": 2 + 6, // rock.paper - win - 2 + 6
		"A Z": 3 + 0, // rock.sciscors - loss - 3 + 0

		"B X": 1 + 0, // paper . rock - loss - 1 + 0
		"B Y": 2 + 3, // paper.paper - draw - 2 + 3
		"B Z": 3 + 6, // paper.scissors - win - 3 + 6

		"C X": 1 + 6, // scissors.rock - win - 1 + 6
		"C Y": 2 + 0, // scissors.paper - loss - 2 + 0
		"C Z": 3 + 3, // scissors.scissors - draw - 3 + 3
	}

	scoresNeedTo := map[string]int{
		"A X": 3 + 0, // rock.lose - play scissors - 3+0
		"A Y": 1 + 3, // rock.draw - play rock -  1 + 3
		"A Z": 2 + 6, // rock.win - play paper - 2 + 6

		"B X": 1 + 0, // paper.lose - play rock - 1 + 0
		"B Y": 2 + 3, // paper.draw - play paper - 2 + 3
		"B Z": 3 + 6, // paper.win - play scissors - 3 + 6

		"C X": 2 + 0, // scissors.lose - play paper - 2 + 0
		"C Y": 3 + 3, // scissors.draw - play scissors - 3 + 3
		"C Z": 1 + 6, // scissors.win - play rock - 1 + 6
	}

	totalScore1 := 0
	totalScore2 := 0

	checkFile, _ := os.Stat("./input.txt")

	if checkFile.Size() == 0 {
		return 0, 0
	}

	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		totalScore1 += scores[text]
		totalScore2 += scoresNeedTo[text]
	}

	return totalScore1, totalScore2
}

func main() {
	score1, score2 := readFile()

	fmt.Printf("total score1: %d, total score2: %d\n", score1, score2)
}
