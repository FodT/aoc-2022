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

func nextRight(r int, c int, rMax int, cMax int) (int, int, bool) {
	return r, c + 1, c < cMax-1
}

func nextLeft(r int, c int, rMax int, cMax int) (int, int, bool) {
	return r, c - 1, c > 0
}

func nextDown(r int, c int, rMax int, cMax int) (int, int, bool) {
	return r + 1, c, r < rMax-1
}

func nextUp(r int, c int, rMax int, cMax int) (int, int, bool) {
	return r - 1, c, r > 0
}

func countFrom(grid [][]Tree, i int, j int, iter func(int, int, int, int) (int, int, bool)) int {
	count := 0
	refHeight := grid[i][j].Height
	maxCols := len(grid[0])
	maxRows := len(grid)
	for r, c, cont := iter(i, j, maxRows, maxCols); cont; r, c, cont = iter(r, c, maxRows, maxCols) {
		count++
		if grid[r][c].Height >= refHeight {
			break
		}
	}
	return count
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

			score *= countFrom(grid, r, c, nextRight)
			if score == 0 {
				continue
			}

			score *= countFrom(grid, r, c, nextLeft)
			if score == 0 {
				continue
			}

			score *= countFrom(grid, r, c, nextDown)
			if score == 0 {
				continue
			}

			score *= countFrom(grid, r, c, nextUp)
			if score == 0 {
				continue
			}

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
