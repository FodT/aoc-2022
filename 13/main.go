package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
		if compare(packets[i], packets[i+1]) {
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

	//const divPackets = `[[2]]
	//[[6]]`
	//lines := strings.Split(divPackets, "\n")
	//for _, p := range lines {
	//	var arr []interface{}
	//	_ = json.Unmarshal([]byte(p), &arr)
	//	packets = append(packets, arr)
	//}
	//
	//sort.SliceStable(packets, func(i, j int) bool {
	//	return compare(packets[i], packets[j])
	//})
	//
	//div1 := 0
	//div2 := 0
	//for i, packet := range packets {
	//	fmt.Println(packet)
	//	s, _ := json.Marshal(packet)
	//	switch string(s) {
	//	case "[[2]]":
	//		div1 = i + 1
	//	case "[[6]]":
	//		div2 = i + 1
	//	}
	//}
	//fmt.Printf("div indexes: %d, %d\n", div1, div2)
	//fmt.Printf("pt2 answer: %d\n", div1*div2)

}

func compare(a, b []interface{}) bool {
	fmt.Printf("comparing %v with %v\n", a, b)

	for ok := len(a) > 0 && len(b) > 0; ok; ok = len(a) > 0 && len(b) > 0 {
		x := a[0]
		a = a[1:]
		y := b[0]
		b = b[1:]

		if c1 := cmp(x, y); c1 != 0 {
			fmt.Printf("COMPARE RETURNS %d\n", c1)
			return c1 < 0
		}

	}

	return len(a) == 0
}

func cmp(x, y interface{}) int {
	typeOfX := reflect.TypeOf(x).Kind()
	typeOfY := reflect.TypeOf(y).Kind()

	if typeOfX == typeOfY && typeOfX == reflect.Float64 {
		if x.(float64) < y.(float64) {
			fmt.Printf("%v is less than %v, returning true\n", x, y)
			return -1
		} else if x.(float64) > y.(float64) {
			fmt.Printf("%v is greater than %v, returning false\n", x, y)
			return 1
		} else {
			fmt.Printf("%v is the same as %v, returning zero\n", x, y)
			return 0
		}
	} else if typeOfX == typeOfY && typeOfX == reflect.Slice {
		c := compare(x.([]interface{}), y.([]interface{}))
		if c {
			return -1
		} else {
			return 1
		}
	} else {
		var newX, newY []interface{}
		if typeOfX == reflect.Float64 {
			fmt.Printf("mixed tpes: converting left %v\n", x)
			newX = append(newX, x)
			newY = y.([]interface{})
		} else {
			fmt.Printf("mixed tpes: converting right %v\n", y)
			newX = x.([]interface{})
			newY = append(newY, y)
		}
		c := compare(newX, newY)
		if c {
			return -1
		} else {
			return 1
		}
	}
}
