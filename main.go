package main

import (
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				log.Printf("[conn: local=%s remote=%s] closing\n", conn.LocalAddr(), conn.RemoteAddr())
				conn.Close()
				break
			} else {
				log.Fatal(err)
			}
		}
		log.Printf("[conn: local=%s remote=%s] read: %d bytes\n", conn.LocalAddr(), conn.RemoteAddr(), n)
		conn.Write(buffer)
	}
}

func main() {
	log.Println("starting listener")
	listen, err := net.Listen("tcp", ""+":"+"8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("handling conns")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[conn: local=%s remote=%s] stablished\n", conn.LocalAddr(), conn.RemoteAddr())
		go handleConn(conn)
	}
}
