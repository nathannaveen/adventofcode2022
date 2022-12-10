package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	"fmt"
)

type pos struct {
	i, j int
}

func main() {
	file, err := os.Open("DayEight.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	arr := [][]int{}
	visible := map[pos] bool {}

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		row := []int{}
		for _, v := range s {
			val, _ := strconv.Atoi(v)
			row = append(row, val)
		}
		arr = append(arr, row)
	}

	for i := 0; i < len(arr); i++ {
		max := -1
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] > max {
				visible[pos{i, j}] = true
				max = arr[i][j]
			}
		}
		max = -1
		for j := len(arr[0]) - 1; j >= 0; j-- {
			if arr[i][j] > max {
				visible[pos{i, j}] = true
				max = arr[i][j]
			}
		}
	}

	for j := 0; j < len(arr[0]); j++ {
		max := -1
		for i := 0; i < len(arr); i++ {
			if arr[i][j] > max {
				visible[pos{i, j}] = true
				max = arr[i][j]
			}
		}
		max = -1
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i][j] > max {
				visible[pos{i, j}] = true
				max = arr[i][j]
			}
		}
	}

	fmt.Println(len(visible))

	file.Close()
}
