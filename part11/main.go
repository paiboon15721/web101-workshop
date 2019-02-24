// 1. ให้ handle route ดังต่อไปนี้
//		1.1 GET / 		 Content = INDEX
//		1.2 GET /profile Content = PROFILE
//		1.3 GET /about Content = ABOUT
//		1.3 GET /contact Content = CONTACT
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

			// build http protocol and html
			body := fmt.Sprintf(`<!DOCTYPE html>
								<html lang="en">
									<head>
										<title></title>
									</head>
									<body>
										<h1>Method = %s</h1>
										<h1>URI = %s</h1>
										<h1>Content = ???</h1>
									</body>
								</html>`, method, uri)
			io.WriteString(conn, "http/1.1 200 ok\r\n")
			fmt.Fprintf(conn, "content-length: %d\r\n", len(body))
			io.WriteString(conn, "content-type: text/html\r\n")
			io.WriteString(conn, "\r\n")
			io.WriteString(conn, body)
			conn.Close()
		}(conn)
	}
}
