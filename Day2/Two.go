package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Day2/DayTwo.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := map[string]int{
		"A": 1, "B": 2, "C": 3,
	}

	res := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line[2] == 'X' { // lose
			if line[0] == 'A' {
				res += 3
			} else if line[0] == 'B' {
				res += 1
			} else if line[0] == 'C' {
				res += 2
			}
		} else if line[2] == 'Y' { // draw
			res += m[string(line[0])] + 3
		} else if line[2] == 'Z' { // win
			if line[0] == 'A' {
				res += 2
			} else if line[0] == 'B' {
				res += 3
			} else if line[0] == 'C' {
				res += 1
			}
			res += 6
		}
	}

	fmt.Println(res)

	file.Close()
}
