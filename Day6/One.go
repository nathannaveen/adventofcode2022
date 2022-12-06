package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"fmt"
)

func main() {
	file, err := os.Open("DaySix.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	s := make([]string, 0)

	for scanner.Scan() {
		s = strings.Split(scanner.Text(), "")
	}

	m := map[string] int{
		"c" : 2,
		"n": 1,
		"z": 1,
	}

	for i := 4; i < len(s); i++ {
		m[s[i]]++

		m[s[i - 4]]--

		if m[s[i - 4]] == 0 {
			delete(m, s[i - 4])
		}

		if len(m) == 4 {
			fmt.Println(i+1)
			break
		}
	}

	file.Close()
}
