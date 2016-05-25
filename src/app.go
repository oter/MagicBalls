package src

import (
	"fmt"
	"net/http"
)

type App struct {
	cellsArray *CellsArray
}

func (a *App) Start(tcpAddr string, httpAddr string) {
	a.cellsArray = NewCellsArray()
	tcpServer := TcpServer{}
	go tcpServer.Start(tcpAddr, a.onUpdateCells)

	StartHttpServer(httpAddr, a.onHttpRequest)
}

func (a *App) onUpdateCells(cellIndex string, count int) {
	a.cellsArray.UpdateCell(cellIndex, count)
}

func (a *App) onHttpRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Largest cell: " + a.cellsArray.GetLargestCellIndex())
	w.Write([]byte(a.cellsArray.GetLargestCellIndex()))
	a.cellsArray.DumpCells()
}
