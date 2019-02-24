// 1. ให้แสดงรูปภาพผ่านทาง browser โดยใช้รูปที่อยู่บน server ตัวเอง
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	li, _ := net.Listen("tcp", ":1234")

	for {
		conn, _ := li.Accept()

		go func(conn net.Conn) {
			i := 0
			var method string
			var uri string
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)

				if i == 0 {
					words := strings.Fields(ln)
					method = words[0]
					uri = words[1]
				}
				if ln == "" {
					break
				}
				i++
			}

			// handle HTML
			if method == "GET" && uri == "/" {
				io.WriteString(conn, "http/1.1 200 ok\r\n")
				io.WriteString(conn, "content-type: text/html\r\n\r\n")
				io.WriteString(conn, `<img src="/cat.jpg">`)
			}
			conn.Close()
		}(conn)
	}
}
