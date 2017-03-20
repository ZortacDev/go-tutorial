package main

import (
	"fmt"
)

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	x, y := swap("hello", "world")
	fmt.Println(x, y)
}
