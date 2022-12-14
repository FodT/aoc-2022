package main

import (
	"fmt"
	"math"
	"strings"
)

type Material int

const (
	Air Material = iota
	Sand
	Rock
	Floor
	Void
)

type Coordinates struct {
	X int
	Y int
}

type Cave struct {
	Bottom    int
	MinX      int
	MaxX      int
	Structure map[Coordinates]Material
}

func main() {

	cave := Cave{
		Bottom:    0,
		MaxX:      0,
		MinX:      math.MaxInt,
		Structure: make(map[Coordinates]Material),
	}

	in := strings.Split(Input, "\n")

	for _, line := range in {
		coords := strings.Split(line, " -> ")
		for i := 0; i < len(coords)-1; i++ {
			c1 := Coordinates{}
			c2 := Coordinates{}
			fmt.Sscanf(coords[i], "%d,%d", &c1.X, &c1.Y)
			fmt.Sscanf(coords[i+1], "%d,%d", &c2.X, &c2.Y)
			drawLine(c1, c2, Rock, &cave)
		}
	}

	// deep copy cave structure because lol
	cave2 := cave
	cave2.Structure = make(map[Coordinates]Material)
	for k, v := range cave.Structure {
		cave2.Structure[k] = v
	}

	voidOutCave(&cave)

	count1 := 0
	for {
		err := sandfall(Coordinates{500, 0}, &cave)
		if err != nil {
			break
		}
		count1++
	}

	fmt.Printf("p1 sand count1:= %d\n", count1)

	addFloorToCave(&cave2)
	voidOutCave(&cave2)
	count2 := 0
	for {
		err := sandfall(Coordinates{500, 0}, &cave2)
		if err != nil {
			break
		}
		count2++
	}

	fmt.Printf("p2 sand count1:= %d\n", count2)
}

func sandfall(c Coordinates, cave *Cave) error {
	if cave.Structure[c] != Air {
		return fmt.Errorf("tried falling into not air")
	}
	if cave.Structure[Coordinates{c.X, c.Y + 1}] == Void {
		return fmt.Errorf("fell into void")
	}
	if cave.Structure[Coordinates{c.X, c.Y + 1}] == Floor {
		//fmt.Println("settled on floor")
		cave.Structure[c] = Sand
		cave.Structure[Coordinates{c.X - 1, c.Y + 1}] = Floor
		cave.Structure[Coordinates{c.X + 1, c.Y + 1}] = Floor
		return nil
	}
	if cave.Structure[Coordinates{c.X, c.Y + 1}] == Air {
		//fmt.Println("falling down")
		return sandfall(Coordinates{c.X, c.Y + 1}, cave)
	} else {
		// try going left
		if cave.Structure[Coordinates{c.X - 1, c.Y + 1}] == Air {
			//fmt.Println("falling left")
			return sandfall(Coordinates{c.X - 1, c.Y + 1}, cave)
		} else if cave.Structure[Coordinates{c.X + 1, c.Y + 1}] == Air {
			// try going right
			//fmt.Println("falling right")
			return sandfall(Coordinates{c.X + 1, c.Y + 1}, cave)
		} else {
			// can't go anywhere, settling
			//fmt.Printf("settling at %v\n", c)
			cave.Structure[c] = Sand
			return nil
		}
	}
}

func drawLine(a, b Coordinates, m Material, cave *Cave) {
	if m == Rock || m == Floor {
		if b.Y > cave.Bottom {
			cave.Bottom = b.Y + 1
		}
		if Min(a.X, b.X) < cave.MinX {
			cave.MinX = Min(a.X, b.X)
		}
		if Max(a.X, b.X) > cave.MaxX {
			cave.MaxX = Max(a.X, b.X)
		}
	}

	dX := b.X - a.X
	dY := b.Y - a.Y

	incX := 0
	if dX != 0 {
		incX = dX / Abs(dX)
	}
	incY := 0
	if dY != 0 {
		incY = dY / Abs(dY)
	}

	for i := 0; i <= Max(Abs(dX), Abs(dY)); i++ {
		c := Coordinates{a.X + i*incX, a.Y + i*incY}
		//fmt.Printf("setting %v to rock\n", c)
		cave.Structure[c] = m
	}

}

func voidOutCave(cave *Cave) {
	start := Coordinates{cave.MinX - 1, cave.Bottom}
	end := Coordinates{cave.MaxX + 1, cave.Bottom}
	drawLine(start, end, Void, cave)
}

func addFloorToCave(cave *Cave) {
	start := Coordinates{cave.MinX - 1, cave.Bottom + 1}
	end := Coordinates{cave.MaxX + 1, cave.Bottom + 1}
	drawLine(start, end, Floor, cave)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
