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

	n := 14

	m := map[string] int{}

	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	for i := n; i < len(s); i++ {
		m[s[i]]++

		m[s[i - n]]--

		if m[s[i - n]] == 0 {
			delete(m, s[i - n])
		}

		if len(m) == n {
			fmt.Println(i+1)
			break
		}
	}

	file.Close()
}
