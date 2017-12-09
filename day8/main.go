package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var values = map[string]int{}
	var inputReader = bufio.NewScanner(strings.NewReader(input))
	for inputReader.Scan() {
		line := inputReader.Text()
		fields := strings.Fields(line)
		var (
			reg  = fields[0]
			op   = fields[1]
			val  = fields[2]
			reg2 = fields[4]
			cond = fields[5]
			val2 = fields[6]
		)

		if _, ok := values[reg]; !ok {
			values[reg] = 0
		}
		if _, ok := values[reg2]; !ok {
			values[reg2] = 0
		}

		v1 := values[reg2]
		v2, err := strconv.Atoi(val2)
		if err != nil {
			fmt.Println("Unable to parse int %s", val2)
		}

		met := false
		switch cond {
		case "<":
			met = v1 < v2
		case "<=":
			met = v1 <= v2
		case ">=":
			met = v1 >= v2
		case ">":
			met = v1 > v2
		case "==":
			met = v1 == v2
		case "!=":
			met = v1 != v2
		}
		if !met {
			continue
		}

		newval, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Unable to parse int %s", val)
		}

		switch op {
		case "inc":
			values[reg] += newval
		case "dec":
			values[reg] -= newval
		}
	}
	max := 0
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
