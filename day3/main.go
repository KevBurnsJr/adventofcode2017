package main

import (
	"fmt"
	"math"
)
/*
   -y|
     |
-x   |
-----|-----
     |   +x
     |
     |+y
*/

func main() {
	var x = 0
	var y = 0
	var size = 1
	var radius = 0
	for i := 0; i < input; i++ {
		if i == size * size {
			// start new spiral
			x += 1
			size += 2
			radius += 1
			continue
		}
		var (
			top    = y == -1 * radius
			left   = x == -1 * radius
			bottom = y == radius
			right  = x == radius
		)
		if top && !left {
			x--
			continue
		}
		if left && !bottom {
			y++
			continue
		}
		if bottom && !right {
			x++
			continue
		}
		if right && !top {
			y--
			continue
		}
	}
	distance := int(math.Abs(float64(x))) + int(math.Abs(float64(y)))
	fmt.Printf("x: %d, y: %d, distance: %d", x, y, distance)
}
