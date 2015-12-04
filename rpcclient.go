// rpcclient
package main

import (
	"github.com/kolonse/TCPServer"
)

func ClientConnect(conn *TCPServer.TCPConn) {
	println("new conn")
}

func ClientDataRecv(conn *TCPServer.TCPConn, data []byte, err error) {
	println("new data")
}
