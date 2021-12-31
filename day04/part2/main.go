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
	boardWin := make([]bool, len(boards))
	for _, n := range nums {
		for b := range boards {
			if boardWin[b] {
				continue
			}
			for i, v := range boards[b] {
				// [1] mark each number
				if v == n {
					boards[b][i] = -1
					break
				}
			}
			// [2] check for a win condition
			// win, err := bingo(boards[b])
			win, err := bingo(boards[b])
			if err != nil {
				log.Fatalln(err)
			}
			if win {
				// [3] sum the answer
				sum := sumUnmarked(boards[b], -1)
				answer := n * sum
				fmt.Println("Answer:", answer)
				boardWin[b] = true
			}
		}
	}
}

// Returns the sum of unmarked numbers on a bingo board
func sumUnmarked(b []int, i int) int {
	s := 0
	for _, n := range b {
		if n != i {
			s += n
		}
	}
	return s
}

// Remove an item from a multi-dimensional slice
func removeItem2D(boards [][]int, i int) [][]int {
	var newBoards [][]int
	for x, b := range boards {
		if x != i {
			newBoards = append(newBoards, b)
		}
	}
	return newBoards
}

// Checks a bingo card to determine if it wins
func bingo(b []int) (bool, error) {
	if len(b) != 25 {
		return false, errors.New(fmt.Sprintf("invalid bingo board length: got %d expected %d", len(b), 25))
	}

	win := true

	// Rows
	steps := []int{0, 5, 10, 15, 20}
	for _, step := range steps {
		start := step
		end := start + 5
		for i := start; i < end; i++ {
			// fmt.Println("Rows: i:", i)
			if b[i] != -1 {
				win = false
			}
		}
		if win {
			return true, nil
		}
		win = true
	}

	win = true

	// Columns
	steps = []int{0, 1, 2, 3, 4}
	for _, step := range steps {
		start := step
		end := 25
		for i := start; i < end; i += 5 {
			// fmt.Println("Cols: i:", i)
			if b[i] != -1 {
				win = false
			}
		}
		if win {
			return true, nil
		}
		win = true
	}

	return false, nil
}
