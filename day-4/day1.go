package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	data := "yzbqklnj"
	num := 0
	for {
		h := data + strconv.Itoa(num)
		hash := md5.Sum([]byte(h))
		hashed := hex.EncodeToString(hash[:])
		
		if hashed[:6] == "000000"{
			fmt.Println(hashed)
			fmt.Println(num)
			return
		}
		num++
	}
}
