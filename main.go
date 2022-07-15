package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func reformat(word string) string {
	var result string
	wr := strings.Replace(word, ":", " ", -1)
	for i := len(wr) - 1; i >= 0; i-- {
		result += string(wr[i])
	}

	return result
}

func main() {

	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		panic(err)
	}

	whatIsIt = reformat(string(sd))

	fmt.Println(whatIsIt)
}
