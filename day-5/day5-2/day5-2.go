package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatalln("ERROR WHILE OPEN THE FILE: ", err)
	}

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		str := scanner.Text()
		if check(str) {
			count++
		}

	}
	fmt.Println(count)
}

func check(str string) bool {
	
	check := true
	l := len(str)
	pairs := make(map[string]int)

	for i := 0; i < l-2; i++ {
		pair := str[i : i+2]
		pairs[pair]++
		//for sandwich "aba"
		if str[i] == str[i+2] {
			check = true
		}
		
	}
	rep := false
	//for pairs that are more than 2
	for _, v := range pairs{
		if v > 1{
			rep = true
			break
		}
	}
	return check && rep
}
