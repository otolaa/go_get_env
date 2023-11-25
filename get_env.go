package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var TOKEN string
var CHAT_ID int

func init() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()

		if strings.Contains(s, "TOKEN=") {
			TOKEN = strings.ReplaceAll(s, "TOKEN=", "")
		}

		if strings.Contains(s, "CHAT_ID=") {
			cid := strings.ReplaceAll(s, "CHAT_ID=", "")
			CHAT_ID, err = strconv.Atoi(cid)
		}
	}
}

func main() {
	// get string id
	fmt.Println(TOKEN)
	fmt.Println(CHAT_ID)
}
