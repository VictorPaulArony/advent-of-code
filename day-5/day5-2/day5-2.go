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
		if check(str) && check2(str) {
			count++
		}

	}
	fmt.Println(count)
}

func check(str string) bool {
	l := len(str)

	for i := 0; i < l-2; i++ {
		// for sandwich "aba"
		if str[i] == str[i+2] {
			return true
		}
	}

	return false
}

func check2(str string) bool {
	l := len(str)
	// for pairs that are more than 2
	for i := 0; i < l-1; i++ {
		for j := i+2;j<l-1;j++{
			if str[j:j+2] == str[i:i+2]{
				return true
			}
		}
		
	}
	return false
}

