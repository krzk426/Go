package main

import (
	"fmt" 
	"strings"
)

func countWordsInString(s string) int {
	return len(strings.Fields(s))
}

func main() {
	var s = "some sentence raz dwa trzy"
	fmt.Println(countWordsInString(s))
}
