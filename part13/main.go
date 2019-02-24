// 1. ให้แสดงรูปภาพผ่านทางหน้า browser
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

			// build http protocol and html
			body := `https://assets.teenvogue.com/photos/5925af0bf5c4720abcde5c0b/3:2/w_1200,h_630,c_limit/cat-fb.jpg`
			io.WriteString(conn, "http/1.1 200 ok\r\n")
			io.WriteString(conn, "content-type: text/html\r\n")
			io.WriteString(conn, "\r\n")
			io.WriteString(conn, body)
			conn.Close()
		}(conn)
	}
}
