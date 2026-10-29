package main

import  "fmt"
import "math/rand/v2"

func main(){
	secret := rand.IntN(100) + 1
	attempts := 0
	
	var guess int

	fmt.Println("=== Welcome to number guesser game ===")
	fmt.Println("Choose number between 1 and 100!")

	for{
		fmt.Print("enter your guess: ")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("enter a valid number bro")
			return
		}

		attempts++

		if guess < secret {
			fmt.Println("higher")
		}else if guess > secret {
			fmt.Println("lower")
		}else{
			fmt.Println("")
			fmt.Println("you are correct")
			fmt.Println("your taken attempts: " , attempts )
			break
		}
	}
}
