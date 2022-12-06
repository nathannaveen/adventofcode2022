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
		"X": 1, "Y": 2, "Z": 3,
	}

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		opp := m[string(line[0])]
		me := m[string(line[2])]

		if opp == me {
			res += me + 3
		} else if opp == 1 && me == 2 {
			res += me + 6
		} else if opp == 1 && me == 3 {
			res += me
		} else if opp == 2 && me == 1 {
			res += me
		} else if opp == 2 && me == 3 {
			res += me + 6
		} else if opp == 3 && me == 1 {
			res += me + 6
		} else if opp == 3 && me == 2 {
			res += me
		}
	}

	fmt.Println(res)

	file.Close()
}
