package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var test bool

// TODO:
// Optimise the bingo function to use a loop
// Complete part 2 of the challenge

func main() {
	flag.BoolVar(&test, "test", false, "run with test input or real input")
	flag.Parse()

	inputFile := "input.txt"
	if test {
		inputFile = "test.txt"
	}

	input, err := os.Open(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer input.Close()

	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// Parse out the bingo numbers
	numsStr := strings.Split(lines[0], ",")
	var nums []int
	for _, v := range numsStr {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		nums = append(nums, i)
	}
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the game boards
	var boards [][]int
	for i := 2; i < len(lines); i += 6 {
		var board []int
		for _, s := range strings.Split(strings.Join(lines[i:i+5], " "), " ") {
			if s == "" {
				continue
			}
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalln(err)
			}
			board = append(board, i)
		}
		boards = append(boards, board)
	}

	// Mark each number [1], check for a win condition [2] and sum the answer [3]
	for _, n := range nums {
		for _, b := range boards {
			for i, v := range b {
				// [1] mark each number
				if v == n {
					b[i] = -1
					break
				}
			}
			// [2] check for a win condition
			win, err := bingo(b)
			if err != nil {
				log.Fatalln(err)
			}
			if win {
				// [3] sum the answer
				sum := 0
				for _, x := range b {
					if x != -1 {
						sum += x
					}
				}
				answer := n * sum
				fmt.Println("Answer:", answer)

				return
			}
		}
	}
}

func bingo(b []int) (bool, error) {
	if len(b) != 25 {
		return false, errors.New("board length not 25, invalid board")
	}

	win := true

	// Check the rows
	for i := 0; i < 5; i++ {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 5; i < 10; i++ {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 10; i < 15; i++ {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 15; i < 20; i++ {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 20; i < 25; i++ {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	// Check the columns
	win = true
	for i := 0; i < 25; i += 5 {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 1; i < 25; i += 5 {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 2; i < 25; i += 5 {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 3; i < 25; i += 5 {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	win = true
	for i := 4; i < 25; i += 5 {
		if b[i] != -1 {
			win = false
		}
	}
	if win {
		return true, nil
	}

	return false, nil
}
