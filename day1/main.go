package main

import (
	"fmt"
	"strconv"
)

func main() {
	var last int
	var total int
	var size = len(input)
	input = input + input
	for i, r := range input {
		n, _ := strconv.Atoi(string(r))
		if n == last {
			total += n
		} else if i > size {
			break
		}
		last = n
	}
	fmt.Println(total)
}
