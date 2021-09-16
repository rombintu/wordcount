package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(len(strings.Split((string(os.Args[1])), " ")))
}
