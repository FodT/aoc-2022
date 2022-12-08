package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Tree struct {
	Height  int
	Visible bool
}

func main() {

	scanner := bufio.NewScanner(strings.NewReader(Input))
	var grid [][]Tree
	for scanner.Scan() {

		line := scanner.Text()
		row := make([]Tree, len(line), len(line))
		for i, c := range line {
			height, _ := strconv.Atoi(string(c))
			row[i] = Tree{
				Height:  height,
				Visible: false,
			}
		}
		grid = append(grid, row)

	}

	fmt.Printf("count visible: %d\n", countVisible(grid))
	fmt.Printf("max scene score: %d\n", sceneScore(grid))
}

func sceneScore(grid [][]Tree) int {
	maxScene := 0
	maxCols := len(grid[0])
	maxRows := len(grid)
	for r := 0; r < maxRows; r++ {
		for c := 0; c < maxCols; c++ {
			score := 1
			refHeight := grid[r][c].Height
			numCanSee := 0
			// search right
			for i := c + 1; i < maxCols; i++ {
				numCanSee++
				if grid[r][i].Height >= refHeight {
					break
				}
			}
			score *= numCanSee
			if score == 0 {
				continue
			}
			numCanSee = 0
			// search left
			for i := c - 1; i >= 0; i-- {
				numCanSee++
				if grid[r][i].Height >= refHeight {
					break
				}
			}
			score *= numCanSee
			if score == 0 {
				continue
			}
			numCanSee = 0
			// search down
			for i := r + 1; i < maxRows; i++ {
				numCanSee++
				if grid[i][c].Height >= refHeight {
					break
				}
			}
			score *= numCanSee
			if score == 0 {
				continue
			}

			numCanSee = 0
			// search up
			for i := r - 1; i >= 0; i-- {
				numCanSee++
				if grid[i][c].Height >= refHeight {
					break
				}
			}
			score *= numCanSee
			if score == 0 {
				continue
			}

			//fmt.Printf("score for %d,%d,: %d\n", r, c, score)
			if score >= maxScene {
				maxScene = score
			}
		}
	}

	return maxScene
}

func countVisible(grid [][]Tree) int {

	maxCols := len(grid[0])
	maxRows := len(grid)

	count := 0

	maxHeight := -1
	for row := 0; row < maxRows; row++ {
		maxHeight = -1
		for col := 0; col < maxCols; col++ {
			if grid[row][col].Height > maxHeight {
				if grid[row][col].Visible != true {
					count++
					grid[row][col].Visible = true
				}
				maxHeight = grid[row][col].Height
			}
		}

		maxHeight = -1
		for col := maxCols - 1; col >= 0; col-- {
			if grid[row][col].Height > maxHeight {
				if grid[row][col].Visible != true {
					count++
					grid[row][col].Visible = true
				}
				maxHeight = grid[row][col].Height
			}
		}

	}

	for col := 0; col < maxCols; col++ {
		maxHeight = -1
		for row := 0; row < maxRows; row++ {
			if grid[row][col].Height > maxHeight {
				if grid[row][col].Visible != true {
					count++
					grid[row][col].Visible = true
				}
				maxHeight = grid[row][col].Height
			}
		}

		maxHeight = -1
		for row := maxRows - 1; row >= 0; row-- {
			if grid[row][col].Height > maxHeight {
				if grid[row][col].Visible != true {
					count++
					grid[row][col].Visible = true
				}
				maxHeight = grid[row][col].Height
			}
		}
	}

	return count
}
