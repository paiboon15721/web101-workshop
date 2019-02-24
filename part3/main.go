// 1. ทำให้ server ปิด connection หลังจากที่ส่งข้อความกลับไปที่ client เรียบร้อยแล้ว
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
	}
}
