package main

import "fmt"

func main() {
	min := 1
	max := 1000000
	fmt.Printf("Think of an integer between %d and %d.\n", min, max)
	fmt.Println("Now I'll try to guess it, using 20 questions or less.")

	var answer string
	for answer != "yes" {
		fmt.Print("Ready? ")
		fmt.Scanf("%s", &answer)
	}

	for i := 0; i < 20; i++ {
		if min == max {
			fmt.Printf("Is it %d? ", min)
			var answer string
			fmt.Scanf("%s", &answer)
			if answer == "yes" {
				fmt.Println("Yay, I win!")
				return
			}
		} else {
			middle := (min + max) / 2
			fmt.Printf("Is it greater than %d? ", middle)
			var answer string
			fmt.Scanf("%s", &answer)
			if answer == "yes" {
				min = middle + 1
			} else {
				max = middle
			}
		}
	}
	if min == max {
		fmt.Printf("Your number must be %d.\n", min)
	} else {
		fmt.Printf("I don't know what it could be. Min is %d, max is %d.\n", min, max)
	}
}
