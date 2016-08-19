package main

import "fmt"

func main() {
	fmt.Println("Welcome to fun game.")
	fmt.Println("You think of a number and I'll try to guess it.")
	for {
		fmt.Print("Is it 7? ")
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == "yes" {
			fmt.Println("Yay, I win!")
			break
		}
	}
}
