package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) == 2 {
		fmt.Println("Size is missing")
		fmt.Println("Usage: [op=*/+-][size]")
		return
	}

	if len(os.Args) != 3 {
		fmt.Println("Usage: [operation = */+ -] [size]")
		return
	}

	size, err := strconv.Atoi(os.Args[2])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	op := os.Args[1]
	switch op {
	case "*":
		multiply := func(i, j int) int {
			return i * j
		}
		head(op, size, multiply)
	case "/":
		divide := func(i, j int) int {
			if j == 0 {
				return 0
			}
			return i / j
		}
		head(op, size, divide)
	case "+":
		add := func(i, j int) int {
			return i + j
		}
		head(op, size, add)
	case "-":
		subtract := func(i, j int) int {
			return i - j
		}
		head(op, size, subtract)
	case "%":
		modulus := func(i, j int) int {
			if j == 0 {
				return 0
			}
			return i % j
		}
		head(op, size, modulus)
	default:
		fmt.Println("Invalid operator.")
		fmt.Println("Valid ops one of: */+-%")
		break
	}

}

func head(op string, size int, opFunc func(int, int) int) {
	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			fmt.Printf("%5d", opFunc(i, j))
		}
		fmt.Println()
	}
	fmt.Println()
}
