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

	curNo := 0
	preNo := 0
	countGreater := 0
	i := 0

	for scanner.Scan() {
		if i == 0 {
			i += 1
			preNo = curNo
			continue
		}

		curNo, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if curNo > preNo {
			countGreater += 1
		}

		preNo = curNo

		i += 1
	}

	fmt.Printf("Answer = %v\n", countGreater)
}
