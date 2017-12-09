package main

import (
	"fmt"
)

func main() {
	sum, pos := scoreGroup(1, 1)
	fmt.Println(pos)
	fmt.Println(sum)
}

func scoreGroup(pos int, depth int) (int, int) {
	var cancelling = false
	var garbage = false
	var score = 0
	for {
		char := input[pos]
		pos++
		if cancelling {
			cancelling = false
			continue
		}
		if garbage {
			if char == '!' {
				cancelling = true
				continue
			}
			if char == '>' {
				garbage = false
			}
			continue
		}
		switch char {
		case '<':
			garbage = true
			continue
		case '{':
			var sub int
			sub, pos = scoreGroup(pos, depth+1)
			score += sub
			continue
		case '}':
			score += depth
			return score, pos
		}
	}
	return score, pos
}
