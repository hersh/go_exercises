package main

import (
	"fmt"
	"os"
)

func main() {
	line := "yes"
	if len(os.Args) > 1 {
		line = ""
		for _, arg := range os.Args[1:] {
			line = line + arg + " "
		}
	}
	for {
		fmt.Println(line)
	}
}
