package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var horizontalPositions []int
	var num int
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		for _, numStr := range line {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			horizontalPositions = append(horizontalPositions, num)
		}
	}
	sort.Ints(horizontalPositions)

	minPos := horizontalPositions[0]
	maxPos := horizontalPositions[len(horizontalPositions)-1]
	var costs []int
	for i := minPos; i < maxPos; i++ {
		cost := 0
		for _, p := range horizontalPositions {
			moves := math.Abs(float64(p - i))
			cost += costOfMoves(moves)
		}
		costs = append(costs, cost)
	}
	sort.Ints(costs)

	fmt.Printf("Answer: %d\n", costs[0])
}

func costOfMoves(moves float64) int {
	c := ((moves * moves) + moves) / 2
	return int(c)
}
