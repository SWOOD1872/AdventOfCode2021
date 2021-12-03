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

	gamma := ""
	epsilon := ""
	n := len(lines[0])
	z, o := 0, 0
	for i := 0; i < n; i++ {
		for _, line := range lines {
			if string(line[i]) == "0" {
				z += 1
			} else {
				o += 1
			}
		}
		if z > o {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
		z = 0
		o = 0
	}
	gammaDec := binToDec(gamma)
	epsilonDec := binToDec(epsilon)

	fmt.Printf("Answer: %v\n", gammaDec*epsilonDec)
}

func binToDec(s string) int {
	a, _ := strconv.ParseInt(s, 2, 64)
	return int(a)
}
