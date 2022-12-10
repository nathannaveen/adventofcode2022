package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type pos struct {
	i, j int
}

func two() {
	file, err := os.Open("Day9/DayNine.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	chain := make([]pos, 10)

	m := map[pos]bool{pos{0, 0}: true}

	for scanner.Scan() {
		line := scanner.Text()

		num, _ := strconv.Atoi(line[2:])
		addI, addJ := 0, 0

		switch line[0] {
		case 'U':
			addI = 1
			addJ = 0
		case 'D':
			addI = -1
			addJ = 0
		case 'L':
			addI = 0
			addJ = -1
		case 'R':
			addI = 0
			addJ = 1
		}

		for i := 0; i < num; i++ {
			chain[0].i += addI
			chain[0].j += addJ

			for j := 1; j < len(chain); j++ {
				H := chain[j-1] // the current head
				T := chain[j]   // the current tail
				if dist(H, T) <= 1 {
					continue
				}

				if H.i != T.i {
					chain[j].i += (H.i - T.i) / abs(H.i-T.i)
				}
				if H.j != T.j {
					chain[j].j += (H.j - T.j) / abs(H.j-T.j)
				}
			}

			m[chain[9]] = true // end of the chain
		}
	}

	fmt.Println(len(m))

	file.Close()
}

func dist(p1, p2 pos) int {
	return max(abs(p1.i-p2.i), abs(p1.j-p2.j))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
