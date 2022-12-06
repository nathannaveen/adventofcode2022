package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day4/DayFour.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()
		elfs := strings.Split(line, ",")
		elf1Range := strings.Split(elfs[0], "-")
		elf2Range := strings.Split(elfs[1], "-")

		a, _ := strconv.Atoi(elf1Range[0])
		b, _ := strconv.Atoi(elf1Range[1])
		c, _ := strconv.Atoi(elf2Range[0])
		d, _ := strconv.Atoi(elf2Range[1])

		if (a <= c && b >= d) || (c <= a && d >= b) ||
			(a <= c && b >= c) || (c <= a && d >= a) {
			res++
		}
	}

	fmt.Println(res)

	file.Close()
}
