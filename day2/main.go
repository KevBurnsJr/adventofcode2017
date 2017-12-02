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
	var sum = 0
	for {
		row, err := inputReader.Read()
		if err != nil {
			break
		}
		high = 0
		low = 0
		for _, field := range row {
			n, _ := strconv.Atoi(string(field))
			if n < low || low == 0 {
				low = n
			}
			if n > high || high == 0 {
				high = n
			}
		}
		sum += high - low
	}
	fmt.Println(sum)
}
