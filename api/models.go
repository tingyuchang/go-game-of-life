package api

import "gameoflife/model"

type reverseParams struct {
	Id     int    `json:"id"`
	Status bool   `json:"is_live"`
	Color  string `json:"color"`
}

type response struct {
	Success bool `json:"success"`
}

type indexResponse struct {
	Cells   []*model.Cell `json:"cells"`
	IsStart bool          `json:"is_start"`
}

type colorResponse struct {
	Color string `json:"color"`
}
