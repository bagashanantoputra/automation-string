package main

import (
	"fmt"
	"strings"
)

func main() {
	data := 
`
6000000
6000001
6000002
6000003
6000004
6000005
`
	dataWithQuotes := addQuotes(data)

	fmt.Println(dataWithQuotes)
}

func addQuotes(input string) string {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		lines[i] = "'" + line + "'" + ","
	}
	return strings.Join(lines, "\n")
}
