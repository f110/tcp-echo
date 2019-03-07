package main

import (
	"log"
	"net"
)

func accept(conn net.Conn) {
	log.Printf("Accept connection from %s", conn.RemoteAddr())
	buf := make([]byte, 1500)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}
		if _, err := conn.Write(buf[:n]); err != nil {
			break
		}
	}
}

func main() {
	l, err := net.Listen("tcp4", ":7000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go accept(conn)
	}
}
