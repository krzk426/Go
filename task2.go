package main

import (
	"fmt"
)

func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return y - x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f())
	}
}


