package main

import (
	"fmt"
	"os"
)

func main() {
	// Example of exit function to terminate the program with a specific exit code. The os.Exit function allows us to exit the program immediately, bypassing any deferred functions.
	fmt.Println("Exiting the program with exit code 1.")
	os.Exit(1) // Uncomment this line to exit the program with code 1
}
