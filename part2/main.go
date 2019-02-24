// 1. ทำให้ server ไม่ปิดตัวเองลง และสามารถ handle request ถัดไปได้
package main

import (
	"io"
	"net"
)

func main() {
	li, _ := net.Listen("tcp", ":1234")
	conn, _ := li.Accept()
	io.WriteString(conn, "Hello from server!")
}
