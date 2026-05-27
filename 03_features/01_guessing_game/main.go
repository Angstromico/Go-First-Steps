package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(10) + 1

	var guess int

	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("Guess a number between 1 and 10.")

	for {
		fmt.Print("Enter your guess: ")

		_, err := fmt.Scan(&guess)

		// ❌ Handle non-numeric input
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")

			// Clear invalid input from buffer
			var discard string
			fmt.Scanln(&discard)

			continue
		}

		// ❌ Handle out-of-range input
		if guess < 1 || guess > 10 {
			fmt.Println("Number must be between 1 and 10. Try again.")
			continue
		}

		// ✅ Valid input → exit loop
		break
	}

	if guess == secretNumber {
		fmt.Println("🎉 Congratulations! You guessed the secret number.")
	} else {
		fmt.Println("❌ Sorry, that's not the secret number.")
		fmt.Println("The secret number was:", secretNumber)
	}
}
