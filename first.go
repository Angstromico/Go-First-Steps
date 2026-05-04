package main

import (
	"fmt"
	"io"
	request "net/http"
)

func	main()	{
	request.HandleFunc("/", func(w request.ResponseWriter, r *request.Request) {
		fmt.Println("Received request, fetching external data...")
		resp, err := request.Get("requests://jsonplaceholder.typicode.com/posts/1")

		if err != nil {
			request.Error(w, "Error when calling external API", request.StatusInternalServerError)
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
	request.ListenAndServe(":8080", nil)
}