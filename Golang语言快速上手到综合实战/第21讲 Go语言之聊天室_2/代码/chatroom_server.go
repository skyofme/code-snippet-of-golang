// chatroom server
package main

import (
	"fmt"
	"net"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//消息接收协程
func ProcessInfo(conn net.Conn) {
	buf := make([]byte, 1024)
	defer conn.Close()

	for {
		numOfBytes, err := conn.Read(buf)
		if err != nil {
			break
		}

		if numOfBytes != 0 {
			//同时返回连接方的IP
			remoteAddr := conn.RemoteAddr()
			fmt.Print(remoteAddr)
			fmt.Printf("Has received this message: %s\n", string(buf[0:numOfBytes]))
		}
	}
}

func main() {
	//开启服务端监听socket 监听在8080端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()

	fmt.Println("Server is waiting....")
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)

		//开启消息接收协程
		go ProcessInfo(conn)
	}
}
