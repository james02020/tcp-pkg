package main

import (
	"encoding/binary"
	"fmt"
	"github.com/matchseller/tcp-pkg"
	"net"
)

const maxBufferSize int = 128
const headerLen int = 6
const lengthOffset int = 4

func serve() {
	tcpAddr := "192.168.1.111:15000"
	listener, err := net.Listen("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	r, err := tcp_package.NewReader(conn, maxBufferSize, headerLen, lengthOffset)
	if err != nil {
		fmt.Println(err)
		return
	}

	go accept(r.Message)

	err = r.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
}

//处理从通道接收的数据
func accept(acceptData chan string) {
	for {
		value, isOk := <-acceptData
		if !isOk {
			break
		}
		parse([]byte(value))
	}
}

func parse(data []byte) {
	//将四字节的包头转为string
	head := string(data[:lengthOffset])
	//将2字节的包体长度转为十进制整形
	size := binary.BigEndian.Uint16(data[lengthOffset : lengthOffset+2])
	//将包体转为string
	body := string(data[lengthOffset+2:])
	fmt.Println(fmt.Sprintf("包头为：%s，包体大小%d字节，值为：%s", head, size, body))
}
