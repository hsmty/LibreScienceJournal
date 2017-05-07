package common

import (
	"fmt"
	"strings"
)

func AskUserInput(msg string) string {
	var input string
	fmt.Print(msg)
	fmt.Scanln(&input)
	return strings.TrimRight(input, "\n")
}
