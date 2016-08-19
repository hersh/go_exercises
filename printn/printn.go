package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Error handling for bonus points. :)
	if len(os.Args) != 2 {
		fmt.Println("USAGE: printn <a positive number>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	// More bonus error handling.
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}
