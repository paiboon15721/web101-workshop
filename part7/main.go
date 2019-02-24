// 1. เพิ่ม http protocol เพื่อให้สามารถทำงานร่วมกับ browser ได้
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

			io.WriteString(conn, "Hello from server!")
			conn.Close()
		}(conn)
	}
}
