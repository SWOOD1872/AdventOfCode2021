package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	var rawData []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		rawData = append(rawData, line)
	}

	var combinations []string
	var encodedOutput []string
	for _, line := range rawData {
		var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 string
		var o1, o2, o3, o4 string
		_, err := fmt.Sscanf(line,
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &o1, &o2, &o3, &o4)
		if err != nil {
			panic(err)
		}
		combinations = append(combinations, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10)
		encodedOutput = append(encodedOutput, o1, o2, o3, o4)
	}

	count := 0
	for _, v := range encodedOutput {
		if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
			count += 1
		}
	}

	fmt.Printf("Answer: %d\n", count)
}
