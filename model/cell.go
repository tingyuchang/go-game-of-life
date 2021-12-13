package model

import "gameoflife/util"

type CellStatus bool

const (
	CELL_DIE  CellStatus = false
	CELL_LIVE CellStatus = true
)

type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Cell struct {
	Coordinate
	Status     CellStatus `json:"is_live"`
	nextStatus CellStatus
	neighbors  []*Cell
	Id         int    `json:"id"`
	Color      string `json:"color"`
}

// CheckLife implements the rule of Conway's Game of Life, please check  https://en.wikipedia.org/wiki/Conway's_Game_of_Life
// each cell store neighbor's pointer, we use these pointers to check cell's life in next step.
func (c *Cell) CheckLife() {
	liveNeighbors := 0
	colors := make([]string, 8)

	for i, v := range c.neighbors {
		if v == nil {
			continue
		}
		if v.Status == CELL_LIVE {
			liveNeighbors++
			colors[i] = v.Color
		}
	}

	if c.Status == CELL_DIE {
		// if exactly 3 live neighbors, cell will get live
		if liveNeighbors == 3 {
			c.nextStatus = CELL_LIVE
			c.Color = util.GetAverageColor(colors)
		}
	} else {
		if liveNeighbors > 3 {
			// if exactly 3 live neighbors, cell will die
			c.nextStatus = CELL_DIE
		} else if liveNeighbors < 2 {
			// if fewer than two live neighbors, cell will die
			c.nextStatus = CELL_DIE
		} else {
			// if two or three live neighbors, cell will live
			c.nextStatus = CELL_LIVE
		}
	}
}

// NextStep updates cell status to next step.
func (c *Cell) NextStep() {
	c.Status = c.nextStatus
}

// SetNeighbors is a setter for cell.neighbors.
func (c *Cell) SetNeighbors(neighbors []*Cell) {
	c.neighbors = neighbors
}

// NewCell initialize a new cell with coordinator x, y and cell status
func NewCell(x, y, id int, status CellStatus) (cell *Cell) {
	return &Cell{Status: status, Coordinate: Coordinate{X: x, Y: y}, Id: id, Color: CELL_COLOR_1}
}
