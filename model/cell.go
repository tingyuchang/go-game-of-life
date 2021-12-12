package model

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
}

// CheckLife implements the rule of Conway's Game of Life, please check  https://en.wikipedia.org/wiki/Conway's_Game_of_Life
// each cell store neighbor's pointer, we use these pointers to check cell's life in next step.
func (c *Cell) CheckLife() {
	liveNeighbors := 0

	for _, v := range c.neighbors {
		if v == nil {
			continue
		}
		if v.Status == CELL_LIVE {
			liveNeighbors++
		}

		if liveNeighbors >= 4 {
			// if more than 4, no need to check
			break
		}
	}

	if c.Status == CELL_DIE {
		// if exactly 3 live neighbors, cell will get live
		if liveNeighbors == 3 {
			c.nextStatus = CELL_LIVE
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
func NewCell(x, y int, status CellStatus) (cell *Cell) {
	return &Cell{Status: status, Coordinate: Coordinate{X: x, Y: y}}
}
