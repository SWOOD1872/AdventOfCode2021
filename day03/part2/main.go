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
		panic(err)
	}
	defer input.Close()

	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	oxygenRating := binToDec(rating(lines, "most"))
	co2Rating := binToDec(rating(lines, "least"))
	lifeSupportRating := oxygenRating * co2Rating
	fmt.Printf("Answer: %v\n", lifeSupportRating)
}

func binToDec(s string) int {
	a, _ := strconv.ParseInt(s, 2, 64)
	return int(a)
}

func rating(list []string, s string) string {
	answer := ""
	n := len(list[0])
	z, o := 0, 0
	for i := 0; i < n; i++ {
		mcn := ""
		for _, line := range list {
			if string(line[i]) == "0" {
				z += 1
			} else {
				o += 1
			}
		}
		if s == "most" {
			if z > o {
				mcn = "0"
			} else {
				mcn = "1"
			}
		} else if s == "least" {
			if z > o {
				mcn = "1"
			} else {
				mcn = "0"
			}
		}
		z = 0
		o = 0
		list = filter(list, i, mcn)
		if len(list) == 1 {
			answer = list[0]
		}
	}
	return answer
}

func filter(list []string, pos int, filter string) []string {
	var filtered []string
	for _, item := range list {
		if string(item[pos]) == filter {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
