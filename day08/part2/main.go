package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("test.txt")
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

	var data []string
	for _, line := range rawData {
		var c1, c2, c3, c4, c5, c6, c7, c8, c9, c10 string
		var o1, o2, o3, o4 string
		_, err := fmt.Sscanf(line,
			"%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&c1, &c2, &c3, &c4, &c5, &c6, &c7, &c8, &c9, &c10, &o1, &o2, &o3, &o4)
		if err != nil {
			panic(err)
		}
		data = append(data, c1, c2, c3, c4, c5, c6, c7, c8, c9, c10, o1, o2, o3, o4)
	}

	allLines := make([][]string, len(data)/14)
	c := 0
	for i := 0; i < len(data); i += 14 {
		var line []string
		if i > len(data)-14 {
			line = data[i:]
		} else {
			line = data[i : i+14]
		}
		allLines[c] = line
		c += 1
	}

	for _, rows := range allLines {
		mapping := make(map[int]string)
		fmt.Println("Rows:\n", rows[0:10])
		for ic, col := range rows[0:10] {
			if ic > 9 {
				break
			}
			// Find the easy numbers
			if len(col) == 2 {
				mapping[1] = col
			}
			if len(col) == 3 {
				mapping[7] = col
			}
			if len(col) == 4 {
				mapping[4] = col
			}
			if len(col) == 7 {
				mapping[8] = col
			}

			// Now find the harder numbers
			if len(col) == 5 {
				if containsAll(col, mapping[7]) {
					mapping[3] = col
					fmt.Printf("%s=%d\n", col, 3)
					continue
				} else if containsN(mapping[4], col) == 2 {
					mapping[2] = col
					fmt.Printf("%s=%d\n", col, 2)
					continue
				} else if containsN(mapping[4], col) == 3 {
					mapping[5] = col
					fmt.Printf("%s=%d\n", col, 5)
					continue
				}
			}
			if len(col) == 6 {
				if containsAll(col, mapping[1]) && containsAll(col, mapping[3]) && containsAll(col, mapping[4]) && containsAll(col, mapping[5]) {
					mapping[9] = col
					fmt.Printf("%s=%d\n", col, 9)
					continue
				} else if containsAll(col, mapping[1]) && containsAll(col, mapping[5]) {
					mapping[6] = col
					fmt.Printf("%s=%d\n", col, 6)
					continue
				} else {
					mapping[0] = col
					fmt.Printf("%s=%d\n", col, 0)
					continue
				}
			}
		}
		fmt.Println(mapping)
	}
}

// mapHasVal returns true if a value is already in a map
func mapHasVal(m map[int]string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

// containsAll returns true if a given string contains all of a given set of characters
func containsAll(s, c string) bool {
	chars := strings.Split(c, "")
	for _, char := range chars {
		if !strings.Contains(s, char) {
			return false
		}
	}
	return true
}

func containsN(from, in string) int {
	fromChars := strings.Split(from, "")
	c := 0
	for _, v := range fromChars {
		if strings.Contains(in, v) {
			c += 1
		}
	}
	return c
}
