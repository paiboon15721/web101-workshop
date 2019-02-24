// 1. เพิ่ม route GET /add เพื่อแสดง form ให้สามารถกดปุ่ม submit เพื่อ POST /add กลับมาที่ server ได้
// 2. เพิ่ม route POST /add ให้ content = ADD DATA SUCCESS
// 3. เพิ่ม link ให้สามารถคลิ๊ก ไป-มา ในแต่ละหน้าได้
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

			var content string
			// router
			if method == "GET" && uri == "/" {
				content = "INDEX"
			}
			if method == "GET" && uri == "/profile" {
				content = "PROFILE"
			}
			if method == "GET" && uri == "/about" {
				content = "ABOUT"
			}
			if method == "GET" && uri == "/contact" {
				content = "CONTACT"
			}

			// build http protocol and html
			body := fmt.Sprintf(`<!DOCTYPE html>
								<html lang="en">
									<head>
										<title></title>
									</head>
									<body>
										<h1>%s</h1>
									</body>
								</html>`, content)
			io.WriteString(conn, "http/1.1 200 ok\r\n")
			fmt.Fprintf(conn, "content-length: %d\r\n", len(body))
			io.WriteString(conn, "content-type: text/html\r\n")
			io.WriteString(conn, "\r\n")
			io.WriteString(conn, body)
			conn.Close()
		}(conn)
	}
}
