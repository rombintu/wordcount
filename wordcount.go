package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	line := strings.Split((string(os.Args[1])), " ")
	fmt.Println(line)
	if line[0] == " " || line[0] == "" {
		fmt.Println(0)
	} else {
		fmt.Println(len(line))
	}
}
