package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var twists = []int{}
	var fields = strings.Split(input, ",")
	for _, field := range fields {
		n, err := strconv.Atoi(field)
		if err != nil {
			fmt.Println("unable to parse int %s", field)
			return
		}
		twists = append(twists, n)
	}

	var circle = []int{}
	for i := 0; i < 256; i++ {
		circle = append(circle, i)
	}

	var pos = 0
	for skip, n := range twists {
		if n > len(circle) {
			fmt.Println("n greater than circle length")
			return
		}
		var segment = []int{}
		if n+pos > len(circle) {
			rem := int(math.Mod(float64(n+pos), float64(len(circle))))
			segment = circle[pos:]
			segment = append(segment, circle[0:rem]...)
			segment = reverse(segment)
			copy(circle[pos:], segment[0:])
			copy(circle[0:], segment[len(segment)-rem:])
		} else {
			segment = circle[pos : pos+n]
			segment = reverse(segment)
			copy(circle[pos:], segment)
		}
		pos = int(math.Mod(float64(pos+n+skip), float64(len(circle))))
	}

	fmt.Printf("Result: %d", circle[0]*circle[1])
}

func reverse(ints []int) []int {
	length := len(ints)
	newslice := make([]int, length)
	for i, v := range ints {
		newslice[length-i-1] = v
	}
	return newslice
}
