package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	go serve()
	go send()
	time.Sleep(30 * time.Second)
}

func send() {
	addr := "192.168.1.111:15000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	head := []byte("head")
	body := []string{
		"I am not afraid of tomorrow for I have seen yesterday and I love today.",
		"If you want to understand today, you have to search yesterday.",
		"You never know what you can do till you try.",
		"A good name keeps its luster in the dark.",
	}
	for {
		var messages []byte
		for _, b := range body {
			body := []byte(b)
			lengthOffset := make([]byte, 2)
			binary.BigEndian.PutUint16(lengthOffset, uint16(len(body)))
			oneMessage := append(head, lengthOffset...)
			oneMessage = append(oneMessage, body...)
			messages = append(messages, oneMessage...)
		}
		_, err := conn.Write(messages)
		if err != nil {
			fmt.Println(err)
			return
		}
		break
	}
}
