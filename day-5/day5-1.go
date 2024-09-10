package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln("ERROR WHILE READING THE FILE:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		str := scanner.Text()
		if check(str){
			count++
		}
	}

	fmt.Println(count)
}

func check(str string) bool {
	pairs := []string{"ab", "cd", "pq", "xy"}
	count := 0
	vol := "aeiou"
	check := false

	for _, v := range pairs{
		if strings.Contains(str, v){
			return false
		}
	}
	for i := 0; i < len(str); i++ {
		for _, v := range vol {
			if string(str[i]) == string(v) {
				count ++
			}
		}
		if i > 0 && str[i] == str[i-1]{
			check = true
		}
	}

	return count >= 3 && check
}
