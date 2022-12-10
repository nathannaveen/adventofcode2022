package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	file, err := os.Open("DayEight.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	arr := [][]int{}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, v := range s {
			val, _ := strconv.Atoi(v)
			row = append(row, val)
		}
		arr = append(arr, row)
	}

	maxScore := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			// each tree
			score := viewingDist(i, j, 0, 1, arr[i][j], 1, arr) *
				viewingDist(i, j, 0, -1, arr[i][j], 1, arr) *
				viewingDist(i, j, 1, 0, arr[i][j], 1, arr) *
				viewingDist(i, j, -1, 0, arr[i][j], 1, arr)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	fmt.Println(maxScore)

	file.Close()
}

func viewingDist(i, j, addI, addJ, height, c int, arr [][]int) int {
	i += addI
	j += addJ
	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr[0]) {
		return 0
	}
	if arr[i][j] >= height {
		return 1
	}
	return 1 + viewingDist(i, j, addI, addJ, height, c + 1, arr)
}