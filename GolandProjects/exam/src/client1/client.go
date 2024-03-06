package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "185.139.70.64:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("send 10\n")
	fmt.Fprintf(conn, "10")
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message from server: " + message)
}
