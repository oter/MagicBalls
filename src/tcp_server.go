package src

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
)

type UpdateCellsFunc func(cellIndex string, count int)

type TcpServer struct {
	onUpdateCells UpdateCellsFunc
	cellsArray    *CellsArray
}

func (ts *TcpServer) Start(laddr string, onUpdateCells UpdateCellsFunc) {
	ts.onUpdateCells = onUpdateCells

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
		ts.handleClientData(data)
	}
}

func (ts *TcpServer) handleClientData(data []byte) {
	packed := bytes.Split(data, []byte(" "))
	if len(packed) != 2 {
		return
	}

	count, err := strconv.Atoi(string(packed[1]))
	if err != nil {
		return
	}
	ts.onUpdateCells(string(packed[0]), count)
}
