package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"
)

type coordinates struct {
	X int
	Y int
}

type dingus struct {
	ind     int
	Coords  coordinates
	Visited map[coordinates]bool
	next    *dingus
}

func newDingus(i int) *dingus {
	return &dingus{
		ind:     i,
		Visited: make(map[coordinates]bool),
	}
}

func (d *dingus) append(din *dingus) {
	d.next = din
}

func (d *dingus) touching(head *dingus) bool {
	diffY := head.Coords.Y - d.Coords.Y
	diffX := head.Coords.X - d.Coords.X
	absX := math.Abs(float64(diffX))
	absY := math.Abs(float64(diffY))

	if absX == 0 && absY == 0 {
		return true
	}
	if absX == 0 && absY == 1 {
		return true
	}
	if absX == 1 && absY == 0 {
		return true
	}
	if absX == 1 && absY == 1 {
		return true
	}

	return false
}

func (d *dingus) follow(head *dingus) {

	if d.touching(head) {
		if d.next == nil {
			d.Visited[d.Coords] = true
			return
		} else {
			d.next.follow(d)
			return
		}
	}

	diffY := head.Coords.Y - d.Coords.Y
	diffX := head.Coords.X - d.Coords.X

	if diffX != 0 && diffY != 0 { //diagonal move
		if math.Abs(float64(diffY)) > 0 {
			if diffY < 0 {
				d.Coords.Y--
			} else {
				d.Coords.Y++
			}
		}
		if math.Abs(float64(diffX)) > 0 {
			if diffX < 0 {
				d.Coords.X--
			} else {
				d.Coords.X++
			}
		}

	} else {
		if math.Abs(float64(diffY)) > 1 {
			if diffY < 0 {
				d.Coords.Y--
			} else {
				d.Coords.Y++
			}
		}

		if math.Abs(float64(diffX)) > 1 {
			if diffX < 0 {
				d.Coords.X--
			} else {
				d.Coords.X++
			}
		}
	}

	d.follow(head)
}

func (d *dingus) countVisited() int {
	count := 0
	for _, v := range d.Visited {
		if v {
			count++
		}
	}
	return count
}

func (d *dingus) move(instr string) {
	var moves int
	incr := 1
	var dir string
	fmt.Sscanf(instr, "%s %d", &dir, &moves)
	if dir == "L" || dir == "D" {
		incr *= -1
	}
	if dir == "U" || dir == "D" {
		for i := 0; i < moves; i++ {
			d.Coords.Y += incr
			if d.next != nil {
				d.next.follow(d)
			}
		}
	}
	if dir == "L" || dir == "R" {
		for i := 0; i < moves; i++ {
			d.Coords.X += incr
			if d.next != nil {
				d.next.follow(d)
			}
		}
	}
}

func main() {
	var firstHead = newDingus(0)
	firstHead.append(newDingus(1))
	firstRopeEnd := firstHead.next

	var secondHead = newDingus(0)

	secondRopeEnd := secondHead
	for i := 1; i < 10; i++ {
		secondRopeEnd.append(newDingus(i))
		secondRopeEnd = secondRopeEnd.next
	}

	scanner := bufio.NewScanner(strings.NewReader(Input))
	for scanner.Scan() {
		firstHead.move(scanner.Text())
		secondHead.move(scanner.Text())
	}
	fmt.Printf("first tail visited %d spots\n", firstRopeEnd.countVisited())
	fmt.Printf("second tail visited %d spots\n", secondRopeEnd.countVisited())
}
