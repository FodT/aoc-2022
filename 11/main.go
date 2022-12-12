package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Item struct {
	WorryScore int
}

type Monkey struct {
	Id              int
	itemsHeld       []Item
	testDenominator int
	falseMonkey     int
	trueMonkey      int
	operationFunc   func(in int) int
	inspectionCount int
}

var monkeys map[int]*Monkey

func (m *Monkey) giveItem(i Item) {
	m.itemsHeld = append(m.itemsHeld, i)
}

func main() {
	monkeys = make(map[int]*Monkey)
	in := strings.Split(Input, "\n\n")
	for _, mIn := range in {
		monkey := parseMonkey(mIn)
		monkeys[monkey.Id] = monkey
	}

	fmt.Println(monkeys)

	rounds := 20

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkeys[j].MonkeyAround()
		}
	}

	for _, m := range monkeys {
		fmt.Printf("monkey %d inspected %d items\n", m.Id, m.inspectionCount)
	}

	fmt.Printf("total monkey business: %d\n", calculateMonkeyBusiness())
}

func (m *Monkey) MonkeyAround() {
	newItems := make([]Item, 0)
	for i, _ := range m.itemsHeld {
		m.inspectionCount++
		m.itemsHeld[i].WorryScore = m.operationFunc(m.itemsHeld[i].WorryScore)
		m.itemsHeld[i].WorryScore /= 3
		if m.itemsHeld[i].WorryScore%m.testDenominator == 0 {
			monkeys[m.trueMonkey].giveItem(m.itemsHeld[i])
		} else {
			monkeys[m.falseMonkey].giveItem(m.itemsHeld[i])
		}
	}
	m.itemsHeld = newItems
}

func calculateMonkeyBusiness() int {
	v := make([]*Monkey, 0, len(monkeys))

	for _, value := range monkeys {
		v = append(v, value)
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i].inspectionCount > v[j].inspectionCount
	})

	return v[0].inspectionCount * v[1].inspectionCount

}

func parseMonkey(monkeyText string) *Monkey {
	id := 0
	monkeyLines := strings.Split(monkeyText, "\n")
	fmt.Sscanf(monkeyLines[0], "Monkey %d:\n", &id)

	itemsStr := strings.TrimPrefix(monkeyLines[1], "  Starting items: ")
	tmp := strings.Split(itemsStr, ", ")
	startingItems := make([]Item, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		startingItems = append(startingItems, Item{v})
	}

	operationLine := strings.TrimPrefix(monkeyLines[2], "  Operation: new = old ")
	op := operationLine[0]
	argStr := operationLine[2:]
	var operation func(int) int

	if op == '+' {
		if argStr == "old" {
			operation = func(in int) int {
				return in + in
			}
		} else {
			v, err := strconv.Atoi(argStr)
			if err != nil {
				log.Fatal(err)
			}
			operation = func(in int) int {
				return in + v
			}
		}
	} else if op == '*' {
		if argStr == "old" {
			operation = func(in int) int {
				return in * in
			}
		} else {
			v, err := strconv.Atoi(argStr)
			if err != nil {
				log.Fatal(err)
			}
			operation = func(in int) int {
				return in * v
			}
		}
	}

	var denom int
	fmt.Sscanf(monkeyLines[3], "  Test: divisible by %d\n", &denom)

	var trueMonkey int
	fmt.Sscanf(monkeyLines[4], "    If true: throw to monkey %d\n", &trueMonkey)

	var falseMonkey int
	fmt.Sscanf(monkeyLines[5], "    If false: throw to monkey %d\n", &falseMonkey)

	return &Monkey{
		Id:              id,
		itemsHeld:       startingItems,
		operationFunc:   operation,
		testDenominator: denom,
		falseMonkey:     falseMonkey,
		trueMonkey:      trueMonkey,
		inspectionCount: 0,
	}
}
