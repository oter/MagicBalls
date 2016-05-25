package main

import (
	. "./src"
)

const (
	TCP_ADDR  = ":9000"
	HTTP_ADDR = ":8000"
)

func main() {
	app := App{}
	app.Start(TCP_ADDR, HTTP_ADDR)
}
