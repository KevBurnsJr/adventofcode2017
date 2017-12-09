package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var validPasswords = 0
	var inputReader = bufio.NewScanner(strings.NewReader(input))
	for inputReader.Scan() {
		line := inputReader.Text()
		fields := strings.Fields(line)
		test := map[string]bool{}
		good := true
		for _, f := range fields {
			_, ok := test[f]
			if ok {
				// map key exists, duplicate word
				good = false
				break
			}
			test[f] = true
		}
		if good {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}
