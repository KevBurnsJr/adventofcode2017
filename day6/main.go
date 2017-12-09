package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var banks = []int{}
	var inputReader = csv.NewReader(strings.NewReader(input))
	inputReader.Comma = '	'
	for {
		row, _ := inputReader.Read()
		for _, field := range row {
			n, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("unable to parse int %s", field)
			}
			banks = append(banks, n)
		}
		break
	}
	cycles := 0
	seen := map[string]bool{}
	for {
		key := seenKey(banks)
		_, ok := seen[key]
		if ok {
			break
		}
		seen[key] = true
		fmt.Println(key)
		i := nextForDistribution(banks)
		banks = distribute(banks, i)
		cycles++
	}
	fmt.Println(cycles)
}

func nextForDistribution(banks []int) int {
	next := 0
	max := 0
	for i, val := range banks {
		if val > max {
			max = val
			next = i
		}
	}
	return next
}

func distribute(banks []int, i int) []int {
	size := len(banks)
	n := banks[i]
	banks[i] = 0
	for j := 1; j <= n; j++ {
		pos := int(math.Mod(float64(i+j), float64(size)))
		banks[pos]++
	}
	return banks
}

func seenKey(banks []int) string {
	stringBank := []string{}
	for _, v := range banks {
		stringBank = append(stringBank, strconv.Itoa(v))
	}
	return strings.Join(stringBank, ",")
}
