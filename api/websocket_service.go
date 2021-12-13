package api

import (
	"encoding/json"
	"gameoflife/model"
	"gameoflife/websocket"
)

func NotifyCellsUpdate() {
	response := indexResponse{
		Cells: model.CurrentController.Cells,
		IsStart: model.CurrentController.IsStart,
	}
	data, _ := json.Marshal(response)
	websocket.CurrentHub.Broadcast(data)
}
