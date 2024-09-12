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
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln("ERROR WHILE READiNG THE FILE: ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	size := 1000
	grid := GridMaker(size)

	var sliced  []string
	for scanner.Scan() {
		lines := scanner.Text()
		sliced = append(sliced, lines)

	}

	for _, options := range sliced {
		Checker(grid, options)
	}

	res := counter(grid)
	fmt.Printf("%d\n", res)
}

func GridMaker(size int) [][]bool {
	grid := make([][]bool, size)

	for i := range grid {
		grid[i] = make([]bool, size)
	}
	return grid
}

func GetParts(s string) (string, string, string) {
	str := strings.Fields(s)
	if str[0] == "toggle"  {
		return "toggle", str[1], str[3]
	}
	return str[0] + " " + str[1], str[2], str[4]
}

func positions(str string) (int, int) {
	part := strings.Split(str, ",")

	x, _ := strconv.Atoi(part[0])
	y, _ := strconv.Atoi(part[1])

	return x, y
}

func Checker(grid [][]bool, str string) {
	signal, start, end := GetParts(str)

	x1, y1 := positions(start)
	x2, y2 := positions(end)

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			switch signal {
			case "turn on":
				grid[i][j] = true
			case "turn off":
				grid[i][j] = false
			case "toggle":
				grid[i][j] = !grid[i][j]
			}
		}
	}
}

func counter(grid [][]bool) int {
	res := 0

	for _, rows := range grid {
		for _, lights := range rows {
			if lights {
				res++
			}
		}
	}

	return res
}
