package main

import (
	"fmt"
	"os"
)

func main() {
	// Example of exit function to terminate the program with a specific exit code. The os.Exit function allows us to exit the program immediately, bypassing any deferred functions.
	fmt.Println("Exiting the program with exit code 1.")
	//Any value different from 0 indicates an error or abnormal termination. By convention, an exit code of 0 indicates successful completion.
	os.Exit(1) // Uncomment this line to exit the program with code 1
	//Exit bypass any defer or panic handling, so the following line will not be executed.
	fmt.Println("This line will not be executed.")
}
