package api

import (
	"context"
	"encoding/json"
	"gameoflife/model"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		response := indexResponse{
			Cells: model.CurrentController.Cells,
			IsStart: model.CurrentController.IsStart,
		}
		data, _ := json.Marshal(response)
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func WsUpgradeHandler(w http.ResponseWriter, r *http.Request) {
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

func StopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

func ReverseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
		model.CurrentController.Reverse(p.Id)
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func NextHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Methods, Content-Type, Authorization")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		model.CurrentController = model.ResetController(model.CurrentController)
		data, _ := json.Marshal(response{Success: true})
		_, _ = w.Write(data)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}