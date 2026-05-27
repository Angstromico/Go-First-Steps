package main

import "fmt"

func main() {
	// Guessing game
	const secretNumber = 7
	var guess int
	println("Welcome to the Guessing Game!")
	println("Try to guess the secret number between 1 and 10.")
	println("Enter your guess:")
	_, err := fmt.Scan(&guess)
	if err != nil {
		println("Error reading input")
		return
	}
	if guess == secretNumber {
		println("Congratulations! You guessed the secret number.")
	} else {
		println("Sorry, that's not the secret number.")
	}
}
