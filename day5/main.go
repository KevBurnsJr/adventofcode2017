package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var lines = []int{}
	var inputReader = bufio.NewScanner(strings.NewReader(input))
	for inputReader.Scan() {
		line := inputReader.Text()
		n, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("unable to parse int %s", line)
		}
		lines = append(lines, n)
	}
	var length = len(lines)
	var steps = 0
	var i = 0
	for i >= 0 && i < length {
		jumps := lines[i]
		lines[i]++
		i += jumps
		steps++
	}
	fmt.Println(steps)
}
