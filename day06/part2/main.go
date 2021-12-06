package main

import (
	"bufio"
	"fmt"
	"os"
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
	var nums []int
	var num int
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, ",")
		for _, numStr := range numsStr {
			num, err = strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
	}

	var bucket = map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, n := range nums {
		bucket[n] += 1
	}

	days := 256
	cycle := 7
	newCycle := 9
	for d := 0; d < days; d++ {
		birthingFish := bucket[0]
		for i := 0; i < newCycle-1; i++ {
			bucket[i] = bucket[i+1]
		}
		bucket[cycle-1] += birthingFish
		bucket[newCycle-1] = birthingFish
	}

	var sums int
	for _, n := range bucket {
		sums += n
	}
	fmt.Printf("Answer: %v", sums)
}
