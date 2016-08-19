package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err == io.EOF {
			break
		}
		fmt.Print(number)
		switch number % 10 {
		case 1:
			fmt.Println("st")
		case 2:
			fmt.Println("nd")
		case 3:
			fmt.Println("rd")
		default:
			fmt.Println("th")
		}
	}
}
