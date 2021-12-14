package main

import (
	"flag"
	"gameoflife/api"
	"gameoflife/model"
	"gameoflife/websocket"
	"net/http"
)

var (
	addr = flag.String("addr", "8080", "http service address")
	size = flag.Int("size", 20, "world size")
)


func main() {
	flag.Parse()
	// websocket init
	websocket.Init()
	go websocket.CurrentHub.Run()
	// init cells and cells controller
	model.Init(*size)
	// register api services
	initAPI()
	// api listener is watching cell controller's each step
	go api.ListenForUpdate()
	err := http.ListenAndServe(":"+*addr, nil)
	if err != nil {
		panic(err)
	}
}

func initAPI() {
	http.Handle("/", http.HandlerFunc(api.IndexHandler))
	http.Handle("/ws", http.HandlerFunc(api.WsUpgradeHandler))
	http.Handle("/start", http.HandlerFunc(api.StartHandler))
	http.Handle("/stop", http.HandlerFunc(api.StopHandler))
	http.Handle("/reverse", http.HandlerFunc(api.ReverseHandler))
	http.Handle("/next", http.HandlerFunc(api.NextHandler))
	http.Handle("/reset", http.HandlerFunc(api.ResetHandler))
	http.Handle("/color", http.HandlerFunc(api.GetColorHandler))
}
