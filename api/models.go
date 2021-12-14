package api

import "gameoflife/model"

// Params for api "/reverse"
type reverseParams struct {
	Id     int    `json:"id"`
	Status bool   `json:"is_live"`
	Color  string `json:"color"`
}

type patternParams struct {
	Pattern     model.CellPattern    `json:"pattern"`
}

// Response for general usage
type response struct {
	Success bool `json:"success"`
}

// Response for api "/"
type indexResponse struct {
	Cells   []*model.Cell `json:"cells"`
	IsStart bool          `json:"is_start"`
}

// Response for api "/color"
type colorResponse struct {
	Color string `json:"color"`
}
