package main

import (
	"fmt"
	"strings"
)

func AskUserInput(msg string) string {
	fmt.Print(msg)
	var input string
	fmt.Scanln(&input)
	return strings.TrimRight(input, "\n")
}
