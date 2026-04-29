# My First Steps in Go

This repository contains my journey as I learn the Go programming language.

## Current Contents

- **Simple HTTP Server** (`main.go`) — A basic TCP HTTP server built from scratch using Go's `net` package. Handles incoming connections with goroutines, parses HTTP requests, and responds with simple HTML. A great introduction to networking and concurrency in Go!

## About This Repo

I'm using this space to practice and experiment with Go fundamentals. As I progress through tutorials, documentation, and exercises, I'll add more programs here covering:

- Basic syntax and data types
- Control flow (loops, conditionals)
- Functions and packages
- Structs and interfaces
- Concurrency (goroutines and channels)
- And more as I learn!

## Running the Code

To run the HTTP server:

```bash
go run main.go
```

Then open your browser and visit: `http://localhost:8080`

Or to build an executable:

```bash
go build -o main.exe main.go
./main.exe
```

---

*This is a learning project — expect frequent changes and experiments as I grow my Go skills!*
