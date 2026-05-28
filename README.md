# My First Steps in Go

<div align="center">
<pre>
 ██████╗  ██████╗ 
██╔════╝ ██╔═══██╗
██║  ███╗██║   ██║
██║   ██║██║   ██║
╚██████╔╝╚██████╔╝
 ╚═════╝  ╚═════╝ 
</pre>
</div>

This repository contains my journey as I learn the Go programming language.

## Folder Structure

```text
.
├── 01_basics/
│   ├── 01_hello/
│   │   └── main.go
│   ├── 02_variables/
│   │   └── main.go
│   └── 03_naming/
│       └── main.go
└── 02_networking/
    ├── 01_tcp_server/
    │   └── main.go
    └── 02_http_server/
        └── main.go
```

## Project Structure

### [01 Basics](./01_basics)
Fundamental concepts of the Go programming language.
- **01 Hello**: The classic "Hello, World!" entry point.
- **02 Variables**: Demonstrates variable declarations, types, short-hand assignments, and scope.
- **03 Naming**: Covers Go naming conventions (camelCase vs. PascalCase) and basic struct usage.

### [02 Networking](./02_networking)
Exploring Go's powerful networking capabilities.
- **01 TCP Server**: A low-level TCP server built with the `net` package that manually parses simple HTTP GET requests.
- **02 HTTP Server**: A higher-level server using the standard `net/http` package that proxies requests to an external API.

## About This Repo

I'm using this space to practice and experiment with Go fundamentals. As I progress through tutorials, documentation, and exercises, I'll add more programs here covering:

- Basic syntax and data types
- Control flow (loops, conditionals)
- Functions and packages
- Structs and interfaces
- Concurrency (goroutines and channels)
- And more as I learn!

## Running the Code

To run any example, navigate to its specific directory and use `go run`:

```bash
cd 01_basics/01_hello
go run main.go
```

Alternatively, you can run it directly from the root of the repository:

```bash
go run 01_basics/01_hello/main.go
```

### Automation Utilities

For even faster execution, you can use the provided scripts from the root. They search for a project by name and run it automatically.

#### Windows (PowerShell)
```powershell
.\run.ps1 hello
.\run.ps1 guessing_game
```

#### Bash (Linux/macOS/Git Bash)
```bash
./run.sh hello
./run.sh guessing_game
```

If multiple projects match your search term, the scripts will list them so you can be more specific.

> **Tip:** You can add an alias to your profile to make this even shorter (e.g., `gr guessing_game`).

For the servers in the `02_networking` section, they will listen on `http://localhost:8080` by default.

---

*This is a learning project — expect frequent changes and experiments as I grow my Go skills!*
