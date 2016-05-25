package src

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

type App struct {
	cellsArray *CellsArray
}

func (a *App) Start(tcpAddr string, httpAddr string) {
	a.cellsArray = NewCellsArray()
	tcpServer := TcpServer{}
	go tcpServer.Start(tcpAddr, a.onTcpReceive)

	StartHttpServer(httpAddr, a.onHttpRequest)
}

func (a *App) onHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Largest cell: " + a.cellsArray.GetLargestCellIndex())
	w.Write([]byte(a.cellsArray.GetLargestCellIndex()))
	a.cellsArray.DumpCells()
}

func (a *App) onTcpReceive(data []byte) {
	packed := bytes.Split(data, []byte(" "))
	if len(packed) != 2 {
		return
	}

	count, err := strconv.Atoi(string(packed[1]))
	if err != nil {
		return
	}
	a.cellsArray.UpdateCell(string(packed[0]), count)
}
