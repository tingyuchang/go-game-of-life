package main

import (
	"gameoflife/api"
	"gameoflife/model"
	"gameoflife/websocket"
	"net/http"
)

var size = 20 // could be input variable

func main() {
	model.CurrentController = model.GetStartWithGlider(size)
	websocket.Init()
	go websocket.CurrentHub.Run()

	http.Handle("/", http.HandlerFunc(api.IndexHandler))
	http.Handle("/ws", http.HandlerFunc(api.WsUpgradeHandler))
	http.Handle("/start", http.HandlerFunc(api.StartHandler))
	http.Handle("/stop", http.HandlerFunc(api.StopHandler))
	http.Handle("/reverse", http.HandlerFunc(api.ReverseHandler))
	http.Handle("/next", http.HandlerFunc(api.NextHandler))
	http.Handle("/reset", http.HandlerFunc(api.ResetHandler))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
