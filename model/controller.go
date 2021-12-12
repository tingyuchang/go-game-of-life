package model

import "fmt"

type Controller struct {
	Cells   []*Cell
	Size    int
	Version int
}

// NewController create a controller which initialize size*size cells
// and setup cell's neighbors.
func NewController(size int) *Controller {
	cells := make([]*Cell, size*size)
	// setup environment
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cell := NewCell(i, j, CELL_DIE)
			cells[i*size+j] = cell
		}
	}

	// set cell's neighbors
	for i, cell := range cells {
		neighbors := make([]*Cell, 8)
		x := i / size
		y := i % size

		if x >= 1 && y >= 1 {
			// left-top
			neighbors[0] = cells[(x-1)*size+y-1]
		}
		if x >= 1 {
			// top
			neighbors[1] = cells[(x-1)*size+y]
		}
		if x >= 1 && y <= size-2 {
			// right-top
			neighbors[2] = cells[(x-1)*size+y+1]
		}
		if y >= 1 {
			// left
			neighbors[3] = cells[x*size+y-1]
		}
		if y <= size-2 {
			// right
			neighbors[4] = cells[x*size+y+1]
		}
		if x <= size-2 && y >= 1 {
			// left-bottom
			neighbors[5] = cells[(x+1)*size+y-1]
		}

		if x <= size-2 {
			// bottom
			neighbors[6] = cells[(x+1)*size+y]
		}

		if x <= size-2 && y <= size-2 {
			// right-bottom
			neighbors[7] = cells[(x+1)*size+y+1]
		}

		cell.SetNeighbors(neighbors)
	}

	return &Controller{Cells: cells, Size: size}
}

// Run execute cell.CheckLife() for all cells
// and then refresh it.
func (c *Controller) Run() {
	// check next step status
	for _, v := range c.Cells {
		v.CheckLife()
	}

	for _, v := range c.Cells {
		v.NextStep()
	}
	c.Version++
}

// Show is a test method, which prints status for all cells
// it could check Run()'s result.
func (c *Controller) Show() {
	fmt.Printf("Step: %v\n", c.Version)
	for i, v := range c.Cells {
		if v.Status {
			fmt.Printf(Yellow("O, "))
		} else {
			fmt.Printf("X, ")
		}
		if (i % c.Size) == 9 {
			fmt.Println()
		}
	}
}

// Reverse reverse cell's status
func (c *Controller) Reverse(position int) {
	c.Cells[position].Status = !c.Cells[position].Status
}
