package model

import (
	"context"
	"fmt"
	"time"
)

type Controller struct {
	Cells   []*Cell
	Size    int
	Version int
	ctx     context.Context
	cancel  context.CancelFunc
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
		if (i % c.Size) == c.Size-1 {
			fmt.Println()
		}
	}
}

// Reverse reverse cell's status
func (c *Controller) Reverse(position int) {
	c.Cells[position].Status = !c.Cells[position].Status
}

// Start starts infinite loop to run c.Run() per second
func (c *Controller) Start(ctx context.Context) {
	if c.cancel != nil {
		return
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	for {
		c.Run()
		c.Show()
		time.Sleep(1 * time.Second)
		select {
		case <-c.ctx.Done():
			return
		default:
		}
	}
}

// Stop use context cancel func to stop infinite loop
func (c *Controller) Stop() {
	if c.cancel != nil {
		c.cancel()
	}
}

// Next likes Run(), but if controller is running
// should stop current loop and run new loop after Run()
func (c *Controller) Next() {
	if c.cancel != nil {
		c.Stop()
		c.Run()
		c.Start(context.Background())
	} else {
		c.Run()
	}
}

// NewController create a controller which initialize size*size cells
// and setup cell's neighbors.
func NewController(size int) *Controller {
	cells := make([]*Cell, size*size)
	// setup environment
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			cell := NewCell(i, j, i*size+j, CELL_DIE)
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

// GetStartWithGlider initialize graph like glider on the center
// ex:
//	X O X
//	X X O
//	O O O
//
func GetStartWithGlider(size int) *Controller {
	controller := NewController(size)
	if size < 10 {
		return controller
	}

	center := (size/2-1)*size + (size/2 - 1)
	controller.Reverse(center - size)
	controller.Reverse(center + 1)
	controller.Reverse(center + size - 1)
	controller.Reverse(center + size)
	controller.Reverse(center + size + 1)

	return controller
}

func ResetController(c *Controller) *Controller {
	if c.cancel != nil {
		c.Stop()
	}
	reset := NewController(c.Size)
	return reset
}
