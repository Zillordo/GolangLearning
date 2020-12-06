package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		_, _ = io.WriteString(conn, "\nHello from TCP server\n")
		_, _ = fmt.Fprintln(conn, "How is your day?")
		_, _ = fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
