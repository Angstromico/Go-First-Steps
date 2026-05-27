package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random generator
	rand.Seed(time.Now().UnixNano())

	// Generate random number between 1 and 10
	secretNumber := rand.Intn(10) + 1

	var guess int
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Try to guess the secret number between 1 and 10.")
	fmt.Print("Enter your guess: ")

	_, err := fmt.Scan(&guess)
	if err != nil {
		fmt.Println("Error reading input")
		return
	}

	if guess == secretNumber {
		fmt.Println("Congratulations! You guessed the secret number.")
	} else {
		fmt.Println("Sorry, that's not the secret number.")
		fmt.Println("The secret number was:", secretNumber) // optional for debugging
	}
}
