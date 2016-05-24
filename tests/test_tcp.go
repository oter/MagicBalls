package main

import (
	"net"
	"bufio"
	"os"
	"fmt"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9000")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
	}
}


