package main

import (
	."./src"
	"net/http"
	"fmt"
)

func main() {
	cellsArray := NewCellsArray()
	tcpServer := TcpServer{}
	go tcpServer.Start(":9000", func(cellIndex string, count int){
		cellsArray.UpdateCell(cellIndex, count)
	})

	httpServer := HttpServer{}
	go httpServer.Start(":8000", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Largest cell: " + cellsArray.GetLargestCellIndex())
		w.Write([]byte(cellsArray.GetLargestCellIndex()))
	})

	select{}
}