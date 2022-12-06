package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("Day1/DayOne.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	arr := []int{}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			arr = append(arr, sum)
			sum = 0
		}
		n, _ := strconv.Atoi(line)
		sum += n
	}

	arr = append(arr, sum)

	sort.Ints(arr)

	fmt.Println(arr[len(arr)-1])

	file.Close()
}
