package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var hasChildren = map[string]bool{}
	var isChild = map[string]bool{}
	var inputReader = bufio.NewScanner(strings.NewReader(input))
	for inputReader.Scan() {
		line := inputReader.Text()
		fields := strings.Fields(line)
		name := fields[0]
		if len(fields) > 3 {
			for i := 3; i < len(fields); i++ {
				isChild[strings.Trim(fields[i], ",")] = true
			}
		} else {
			continue
		}
		hasChildren[name] = true
	}
	matches := []string{}
	for name, _ := range hasChildren {
		_, ok := isChild[name]
		if !ok {
			matches = append(matches, name)
		}
	}
	fmt.Println(matches)
}
