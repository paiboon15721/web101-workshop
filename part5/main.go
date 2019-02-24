// ทำให้ server handle หลายๆ request พร้อมๆ กันได้
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
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		io.WriteString(conn, "Hello from server!")
		conn.Close()
	}
}
