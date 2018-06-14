package main

import (
	"os"
	"fmt"
	"net"
	"bytes"
	"io"
)

/*
无论什么协议简历什么形式连接，都只需要调用net.Dial()
func Dial(net, addr string) (Conn, error)
net:网络协议名字
addr:IP地址或域名，端口号以:的形式在地址或域名后面。如果连接成功，返回连接对象，否则返回error。
*/

/*
TCP链接：
conn, err := net.Dial("tcp", "xxxx:xxx")
UDP链接：
conn, err := net.Dial("udp", "xxxx:xxx")
ICMP链接：协议名称
conn, err := net.Dial("ip4:icmp", "xxxx")
ICMP链接：协议编号
conn, err := net.Dial("ip4:1", "xxxx")
目前支持：tcp,tcp4,tcp6,udp,udp4,udp6,ip,ip4,ip6
使用conn的Write()发送数据，Read()方法接收数据。
*/

func main() {
	if len(os.Args) != 2 {
		fmt.Println("输入不正确")
		os.Exit(1)
	}

	service := os.Args[1]
	conn, err := net.Dial("ip4:icmp", service)
	checkError(err)

	var msg [512]byte
	msg[0] = 8 //echo
	msg[1] = 0 //code 0
	msg[2] = 0 //checksum
	msg[3] = 0 //checksum
	msg[4] = 0 //identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0 //sequence[0]
	msg[7] = 37 //sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check>>8)
	msg[3] = byte(check&255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("反馈结果")
	if msg[5] == 13 {
		fmt.Println("identifier通过")
	}
	if msg[7] == 37 {
		fmt.Println("sequence通过")
	}
	os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n:=1; n<len(msg)-1; n+=2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum>>16) + (sum&0xffff)
	sum += (sum>>16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Println("发生了错误：",err.Error())
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

