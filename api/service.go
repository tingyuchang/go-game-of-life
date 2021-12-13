package api

import (
	"encoding/json"
	"gameoflife/model"
	"gameoflife/websocket"
)

// ListenForUpdate listen controller's each step update and
// send broadcast message to each client
// TODO: stop channel receiver, and improve performance
func ListenForUpdate() {
	for {
		select {
		case <- model.CurrentController.Step:
			broadcastUpdate()
		default:
		}
	}
}

func broadcastUpdate() {
	response := indexResponse{
		Cells: model.CurrentController.Cells,
		IsStart: model.CurrentController.IsStart,
	}
	data, _ := json.Marshal(response)
	websocket.CurrentHub.BroadcastMsg(data)
}
