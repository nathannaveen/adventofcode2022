package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("Day3/DayThree.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := 0

	m := map[string]bool{}
	m2 := map[string]bool{}

	c := 0

	for scanner.Scan() {
		line := scanner.Text()
		c++

		for _, letter := range line {
			if c%3 == 0 && m2[string(letter)] && m[string(letter)] {
				res += priority2(letter)
				m = map[string]bool{}
				m2 = map[string]bool{}
			} else if c%3 == 1 {
				m[string(letter)] = true
			} else if c%3 == 2 {
				m2[string(letter)] = true
			}
		}

	}

	fmt.Println(res)

	file.Close()
}

func priority2(letter rune) int {
	if unicode.IsUpper(letter) {
		return 27 + int(letter-'A')
	} else {
		return 1 + int(letter-'a')
	}
}
