package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

var (
	id = 1
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go func(id int, conn net.Conn) {
			defer func() {
				if err := conn.Close(); err != nil {
					fmt.Println(time.Now(), id, "close error", err)
				} else {
					fmt.Println(time.Now(), id, "close success")
				}
			}()

			// 1
			if writeLen, err := conn.Write([]byte("Hello World!")); err != nil {
				fmt.Println(time.Now(), id, "write hello world error", err)
				return
			} else {
				fmt.Println(time.Now(), id, "write hello world success", writeLen)
			}

			// 2
			_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
			buf := make([]byte, 64*1024)
			for {
				if readLen, err := conn.Read(buf); err != nil {
					fmt.Println(time.Now(), id, "read error", err)
					return
				} else {
					fmt.Println(time.Now(), id, "read success", string(buf[:readLen]))
					_ = conn.SetReadDeadline(time.Now().Add(3 * time.Second))
				}
			}

		}(id, conn)
		id++
	}
}
