package main

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var inputReader = csv.NewReader(strings.NewReader(input))
	inputReader.Comma = '	'
	var high int
	var low int
	var reset bool
	var sum = 0
	for {
		row, err := inputReader.Read()
		if err != nil {
			break
		}
		reset = true
		for _, field := range row {
			n, _ := strconv.Atoi(field)
			if n < low || reset {
				low = n
			}
			if n > high || reset {
				high = n
			}
			reset = false
		}
		sum += high - low
	}
	fmt.Println(sum)
}
