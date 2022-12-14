package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {

	in := strings.Split(Input, "\n\n")

	var packets [][]interface{}
	for _, linePair := range in {
		lines := strings.Split(linePair, "\n")
		for _, line := range lines {
			var arr []interface{}
			_ = json.Unmarshal([]byte(line), &arr)
			//fmt.Printf("Unmarshaled: %v\n", arr)
			packets = append(packets, arr)
		}
	}

	//fmt.Println(packets)

	var pairIndices []int
	for i, pair := 0, 1; i < len(packets); i, pair = i+2, pair+1 {
		//fmt.Printf("******** comparing %v to %v *********\n", packets[i], packets[i+1])
		if cmp2(packets[i], packets[i+1]) <= 0 {
			fmt.Printf("%d is good!\n", pair)
			pairIndices = append(pairIndices, pair)
		}
	}
	fmt.Printf("pairIndices: %v\n", pairIndices)
	sum := 0
	for _, x := range pairIndices {
		sum += x
	}
	fmt.Printf("sum: %v\n", sum)

	const divPackets = `[[2]]
	[[6]]`
	lines := strings.Split(divPackets, "\n")
	for _, p := range lines {
		var arr []interface{}
		_ = json.Unmarshal([]byte(p), &arr)
		packets = append(packets, arr)
	}

	sort.SliceStable(packets, func(i, j int) bool {
		return cmp2(packets[i], packets[j]) < 0
	})

	div1 := 0
	div2 := 0
	for i, packet := range packets {
		fmt.Println(packet)
		s, _ := json.Marshal(packet)
		switch string(s) {
		case "[[2]]":
			div1 = i + 1
		case "[[6]]":
			div2 = i + 1
		}
	}
	fmt.Printf("div indexes: %d, %d\n", div1, div2)
	fmt.Printf("pt2 answer: %d\n", div1*div2)

}

func cmp2(x, y any) int {
	xv, xok := x.([]any)
	yv, yok := y.([]any)

	switch {
	case !xok && !yok:
		return int(x.(float64) - y.(float64))
	case !xok:
		xv = []any{x}
	case !yok:
		yv = []any{y}
	}

	for i := 0; i < len(xv) && i < len(yv); i++ {
		if comp := cmp2(xv[i], yv[i]); comp != 0 {
			return comp
		}
	}

	return len(xv) - len(yv)
}
