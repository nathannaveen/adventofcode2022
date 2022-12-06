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

	for scanner.Scan() {
		line := scanner.Text()

		m := map[string]bool{}

		for _, letter := range line[:len(line)/2] {
			m[string(letter)] = true
		}

		for _, letter := range line[len(line)/2:] {
			if m[string(letter)] {
				res += priority(letter)
				break
			}
		}
	}

	fmt.Println(res)

	file.Close()
}

func priority(letter rune) int {
	if unicode.IsUpper(letter) {
		return 27 + int(letter-'A')
	} else {
		return 1 + int(letter-'a')
	}
}
