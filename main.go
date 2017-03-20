package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Bitte gib ein bis zu welcher Grenze Fibonacci-Zahlen ausgegeben werden sollen:")
	var input string
	fmt.Scanln(&input)

	i, err := strconv.Atoi(input)
	for err != nil {
		fmt.Println("Gib gefälligst ne Zahl ein!")
		fmt.Scanln(&input)
		i, err = strconv.Atoi(input)
	}
	x := 1
	y := 1
	fmt.Println(x)
	for y <= i {
		fmt.Println(y)
		z := x + y
		x = y
		y = z
	}

}
