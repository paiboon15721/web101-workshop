// 1. แก้ไข และส่งเป็น html แทน เพื่อให้ browser สามารถแสดงเป็นตัวหนาได้
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	li, _ := net.Listen("tcp", ":1234")

	for {
		conn, _ := li.Accept()

		go func(conn net.Conn) {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)

				if ln == "" {
					break
				}
			}

			io.WriteString(conn, "http/1.1 200 ok\r\n")
			io.WriteString(conn, "content-type: text/plain\r\n")
			io.WriteString(conn, "\r\n")

			io.WriteString(conn, "Hello from server!")
			conn.Close()
		}(conn)
	}
}
