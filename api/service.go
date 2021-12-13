package api

import (
	"encoding/json"
	"gameoflife/model"
	"gameoflife/websocket"
	"math/rand"
)

// ListenForUpdate listen controller's each step update and
// send broadcast message to each client
// TODO: stop channel receiver, and improve performance
func ListenForUpdate() {
	for {
		<-model.CurrentController.Step
		broadcastUpdate()
	}
}

// broadcastUpdate send formatted message to websocket hub
func broadcastUpdate() {
	response := indexResponse{
		Cells:   model.CurrentController.Cells,
		IsStart: model.CurrentController.IsStart,
	}
	data, _ := json.Marshal(response)
	websocket.CurrentHub.BroadcastMsg(data)
}

// randomColor returns a random color from pre-defined colors
func randomColor() string {
	colors := []string{model.CELL_COLOR_1, model.CELL_COLOR_2, model.CELL_COLOR_3, model.CELL_COLOR_4, model.CELL_COLOR_5, model.CELL_COLOR_6, model.CELL_COLOR_7, model.CELL_COLOR_8, model.CELL_COLOR_9, model.CELL_COLOR_10, model.CELL_COLOR_11, model.CELL_COLOR_12, model.CELL_COLOR_13, model.CELL_COLOR_14, model.CELL_COLOR_15}

	r := rand.Intn(len(colors))

	return colors[r]
}
