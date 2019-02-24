// 1. อ่านค่า method และ uri ที่ส่งมาจาก client เก็บใส่ตัวแปร
// 2. พิมพ์ออกมาทาง console
// 3. แสดงผลผ่านทาง html กลับไปที่ browser
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

			body := `<!DOCTYPE html>
								<html lang="en">
									<head>
										<title></title>
									</head>
									<body>
										<h1>Method = ???</h1>
										<h1>URI = ???</h1>
									</body>
								</html>`
			io.WriteString(conn, "http/1.1 200 ok\r\n")
			fmt.Fprintf(conn, "content-length: %d\r\n", len(body))
			io.WriteString(conn, "content-type: text/html\r\n")
			io.WriteString(conn, "\r\n")
			io.WriteString(conn, body)
			conn.Close()
		}(conn)
	}
}
