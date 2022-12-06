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
	file, err := os.Open("Day5/DayFive.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	stacks := make([][]string, 9)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 && line[1] == '1' { //  end of stacks
			break
		}

		for i := 0; i < len(line); i++ {
			if i*4 < len(line) && line[i*4] != ' ' {
				stacks[i] = append([]string{string(line[i*4+1])}, stacks[i]...)
			}
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		moves := strings.Split(line, " ")

		n, _ := strconv.Atoi(moves[1])
		from, _ := strconv.Atoi(moves[3])
		to, _ := strconv.Atoi(moves[5])

		from -= 1
		to -= 1

		stacks[from], stacks[to] =
			stacks[from][:len(stacks[from])-n], append(stacks[to], stacks[from][len(stacks[from])-n:]...)
	}

	res := ""

	for _, stack := range stacks {
		res += stack[len(stack)-1]
	}

	fmt.Println(res)

	file.Close()
}
