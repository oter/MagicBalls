package src

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type TcpReceiveFunc func(data []byte)

type TcpServer struct {
	onTcpReceive TcpReceiveFunc
	cellsArray   *CellsArray
}

func (ts *TcpServer) Start(laddr string, onTcpReceive TcpReceiveFunc) {
	ts.onTcpReceive = onTcpReceive

	listener, _ := net.Listen("tcp", laddr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error while accepting a TCP connection: " + err.Error())
		} else {
			go ts.handleClient(conn)
		}
	}
}

func (ts *TcpServer) handleClient(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	for {
		data, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Got error: " + err.Error())
			return
		}
		ts.onTcpReceive(data)
	}
}
