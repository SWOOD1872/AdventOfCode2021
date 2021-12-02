package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *position) move(n int) {
	p.horizontal += n
	p.depth += p.aim * n
}

func (p *position) climb(n int) {
	p.aim -= n
}

func (p *position) dive(n int) {
	p.aim += n
}

func main() {
	p := position{0, 0, 0}

	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		i := line[0]
		v, _ := strconv.Atoi(line[1])

		if strings.ToLower(i) == "forward" {
			p.move(v)
		}
		if strings.ToLower(i) == "down" {
			p.dive(v)
		}
		if strings.ToLower(i) == "up" {
			p.climb(v)
		}
	}

	answer := p.horizontal * p.depth
	fmt.Printf("Answer: %v\n", answer)
}
