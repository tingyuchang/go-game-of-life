package api

import (
	"context"
	"encoding/json"
	"gameoflife/model"
	"gameoflife/websocket"
	"net/http"
)

// using go-gin (web framework) is more convenient
// TODO: change api to go-gin

// IndexHandler returns all cells and controller's status
// TODO: is need always return all cells or only updated cells?
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		response := indexResponse{
			Cells:   model.CurrentController.Cells,
			IsStart: model.CurrentController.IsStart,
		}
		data, _ := json.Marshal(response)
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// WsUpgradeHandler handles websocket upgrade
func WsUpgradeHandler(w http.ResponseWriter, r *http.Request) {
	websocket.ServeWS(w, r, websocket.CurrentHub)
}

// StartHandler runs cell controller's Start(), which runs each steps automatically
func StartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		go model.CurrentController.Start(context.Background())
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// StopHandler stops cell controller's automation
func StopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		model.CurrentController.Stop()
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ReverseHandler makes cell from live to die and vice
// it receives color from client, and assign to reversed cell
func ReverseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		var p reverseParams
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Params are invalid", http.StatusBadRequest)
		}
		model.CurrentController.Reverse(p.Id, p.Color)
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// NextHandler executes cell controller's Next()
func NextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		model.CurrentController.Next()
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// ResetHandler executes cell controller's Reset()
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		model.CurrentController.Reset()
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GetColorHandler returns a random color to client
func GetColorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		response := colorResponse{
			Color: randomColor(),
		}
		data, _ := json.Marshal(response)
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// SetPatternHandler set predefine pattern to broad
func SetPatternHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// to resolve CORS issue when local testing
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		var p patternParams
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, "Params are invalid", http.StatusBadRequest)
		}
		model.SetToPatterns(p.Pattern, model.CurrentController.Size)
		response := indexResponse{
			Cells:   model.CurrentController.Cells,
			IsStart: model.CurrentController.IsStart,
		}
		data, _ := json.Marshal(response)
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
