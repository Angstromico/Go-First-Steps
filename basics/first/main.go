package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go func(c net.Conn) {
			defer c.Close()
			scanner := bufio.NewScanner(c)
			for scanner.Scan() {
				req := scanner.Text()
				fmt.Println(req)
				if strings.Contains(req, "GET") {
					c.Write([]byte("HTTP/1.1 200 OK\n\n"))
					c.Write([]byte("<html><body><h1>Hello World!</h1></body></html>"))
				} else {
					c.Write([]byte("HTTP/1.1 404 Not Found\n\n"))
					c.Write([]byte("<html><body><h1>Not Found</h1></body></html>"))
				}
			}
		}(conn)
	}
}
