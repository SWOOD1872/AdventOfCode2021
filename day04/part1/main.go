package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	// inputLen := len(lines)

	var boardLine []int
	// var game [][]int
	// var games []game

	var numbers []int
	i := 0
	c := 0
	for _, line := range lines {
		// Get all the bingo numbers (always the first line in the input)
		if i == 0 {
			lineSplit := strings.Split(line, ",")
			for _, v := range lineSplit {
				number, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number)
			}
			i += 1
			continue
		}

		// Skip the first blank line
		if line == "" && i == 1 {
			i += 1
			continue
		}

		// Parse out all the game boards
		lineSplit := removeWhitespace(strings.Split(strings.TrimSpace(line), " "))
		for _, v := range lineSplit {
			num, err := strconv.Atoi(strings.TrimSpace(v))
			if err != nil {
				panic(err)
			}
			boardLine = append(boardLine, num)
			fmt.Printf("Board Line: %v\n", boardLine)
		}
		if len(boardLine) == 5 {
			// append(g.board[c][], boardLine)
			c += 1
		}

		i += 1
	}
}

func removeWhitespace(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			newLines = append(newLines, line)
		}
	}
	return newLines
}
