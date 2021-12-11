package model

type CellStatus bool

const (
	CELL_DIE  CellStatus = false
	CELL_LIVE CellStatus = true
)

type Coordinate struct {
	X int
	Y int
}

type Cell struct {
	Coordinate
	Status     CellStatus `json:"is_live"`
	nextStatus CellStatus
	neighbors  []*Cell
}

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

func (c *Cell) SetNeighbors(neighbors []*Cell) {
	c.neighbors = neighbors
}

func (c *Cell) Refresh() {
	c.Status = c.nextStatus
}

func NewCell(x, y int, status CellStatus) (cell *Cell) {
	return &Cell{Status: status, Coordinate: Coordinate{X: x, Y: y}}
}