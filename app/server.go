package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	// Uncomment this block to pass the first stage
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		defer conn.Close()
		
		func(conn net.Conn) {
			defer conn.Close()

			for {
				req := make([]byte, 1024)
				n, err := conn.Read(req)
				if err == io.EOF {
					fmt.Println("EOF")
					return
				}
				if err != nil {
					fmt.Println("Error reading from connection: " , err.Error())
					os.Exit(1)
				}
				fmt.Printf("Read %d bytes from connection: %s\n", n, string(req))

				toks := strings.Split(string(req), "\r\n")
				fmt.Printf("toks: %v\n", toks)

				var length int
				fmt.Sscanf(toks[0],"*%d", &length)
				fmt.Printf("length: %d\n", length)

				n, err = conn.Write([]byte("+PONG\r\n"))
				if err != nil {
					fmt.Println("Error writing to connection: " , err.Error())
					os.Exit(1)
				}
				fmt.Printf("Wrote %d bytes to connection\n", n)
			}
		}(conn)
	}
}
