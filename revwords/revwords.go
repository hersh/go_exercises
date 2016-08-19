package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		words := strings.Fields(line)
		for i := len(words) - 1; i >= 0; i-- {
			fmt.Print(words[i], " ")
		}
		fmt.Println()
		if err == io.EOF {
			break
		}
	}
}
