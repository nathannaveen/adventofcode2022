package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func one() {
	file, err := os.Open("Day9/DayNine.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	H := pos{0, 0}
	T := pos{0, 0}

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
			H.i += addI
			H.j += addJ
			if dist(H, T) <= 1 {
				continue
			}

			if H.i != T.i {
				T.i += (H.i - T.i) / abs(H.i-T.i)
			}
			if H.j != T.j {
				T.j += (H.j - T.j) / abs(H.j-T.j)
			}

			m[T] = true
		}
	}

	fmt.Println(len(m))

	file.Close()
}
