package api

type reverseParams struct {
	Id     int  `json:"id"`
	Status bool `json:"is_live"`
}

type response struct {
	Success bool `json:"success"`
}

