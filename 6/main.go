package main

import "fmt"

func recurringChar(str string) bool {
	recurring := map[rune]int{}

	for _, char := range str {
		if _, ok := recurring[char]; ok {
			return true
		} else {
			recurring[char] = 1
		}
	}

	return false
}

func main() {
	for i := 0; i < len(Input)-3; i++ {
		if !recurringChar(Input[i : i+4]) {
			fmt.Printf("NOT RECURRING AT i=%d\n", i)

			fmt.Printf("marker after character %d\n", i+4)
			break
		}
		//fmt.Println(Input[i : i+4])
	}

	for i := 0; i < len(Input)-13; i++ {
		if !recurringChar(Input[i : i+14]) {
			fmt.Printf("NOT RECURRING AT i=%d\n", i)

			fmt.Printf("marker after character %d\n", i+14)
			break
		}
		//fmt.Println(Input[i : i+4])
	}
}
