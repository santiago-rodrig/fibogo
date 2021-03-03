package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var errInvalidArgument = errors.New("Invalid arguments")

func main() {
	if len(os.Args) != 2 {
		printHelp()
		os.Exit(1)
	}

	n, err := parseArgument(os.Args[1])

	if err != nil {
		printHelp()
		os.Exit(2)
	}

	fmt.Printf("The number at the position %d in the Fibonacci sequence is %d\n", n, fibo(n))
}

func fibo(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

func printHelp() {
	fmt.Println("Please give one number greater than, or equal to 1")
	fmt.Println("Example: fibogo 5")
}

func parseArgument(arg string) (int, error) {
	n, err := strconv.ParseInt(arg, 10, 64)

	if err != nil {
		return 0, errInvalidArgument
	}

	if n < 1 {
		return 0, errInvalidArgument
	}

	return int(n), nil
}
