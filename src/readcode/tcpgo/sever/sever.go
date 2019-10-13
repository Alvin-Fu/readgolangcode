package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	tcpListen, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListen.Close()
	for {
		tcpConn, err := tcpListen.AcceptTCP()
		if err != nil {
			return
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnected :" + ipStr)
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(string(message))
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
