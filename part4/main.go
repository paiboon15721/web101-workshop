// 1. อ่านข้อความที่ได้รับจาก client
// 2. พิมพ์ข้อความออกมาทาง console ทีละบันทัด
package main

import (
	"io"
	"net"
)

func main() {
	li, _ := net.Listen("tcp", ":1234")

	for {
		conn, _ := li.Accept()
		io.WriteString(conn, "Hello from server!")
		conn.Close()
	}
}
