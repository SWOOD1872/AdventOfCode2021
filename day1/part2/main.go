package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var n1, n2, n3 int = 0, 0, 0
	countGreater := 0

	// Need the first 3 numbers initially so we can start comparing and shifting them in Scan()
	scanner.Scan()
	n3, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	n2, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	n1, _ = strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		// Get the next number from the input
		n, _ := strconv.Atoi(scanner.Text())

		// If the next/latest number is greater than oldest (n+3), increment the counter
		if n > n3 {
			countGreater++
		}

		// Ensure we shift each number for the next loop
		n3 = n2
		n2 = n1
		n1 = n
	}

	fmt.Printf("Answer: %v\n", countGreater)
}
