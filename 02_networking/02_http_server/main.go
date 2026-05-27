package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request, fetching external data...")
		resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

		if err != nil {
			http.Error(w, "Error when calling external API", http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()

		fmt.Fprintf(w, "Hello World! \n")
		fmt.Fprintf(w, "HTTP Response status: %s \n\n", resp.Status)

		w.Header().Set("Content-Type", "application/json")
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			fmt.Println("Error in copying the response:", err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
