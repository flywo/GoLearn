package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("输入错误")
		os.Exit(1)
	}

	server := os.Args[1]
	conn, err := net.Dial("tcp", server)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(1)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("发生了错误",err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
