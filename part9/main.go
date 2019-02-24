// 1. ทำให้ browser รองรับภาษาไทย
// 2. ทำให้ browser ปิด connection เองได้
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

			// utf-8 byte slice
			body := []byte{224, 184, 151, 224, 184, 148, 224, 184, 170, 224, 184, 173, 224, 184, 154, 224, 184, 160, 224, 184, 178, 224, 184, 169, 224, 184, 178, 224, 185, 132, 224, 184, 151, 224, 184, 162}
			// tis-620 byte slice
			// body := []byte{183, 180, 202, 205, 186, 192, 210, 201, 210, 228, 183, 194}

			io.WriteString(conn, "http/1.1 200 ok\r\n")
			io.WriteString(conn, "content-type: text/plain\r\n")
			io.WriteString(conn, "\r\n")
			io.WriteString(conn, string(body))
			conn.Close()
		}(conn)
	}
}
