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
	IsStart bool
	Step    chan struct{}
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
func (c *Controller) Reverse(position int, color string) {
	c.Cells[position].Status = !c.Cells[position].Status
	if c.Cells[position].Status == CELL_LIVE {
		if len(color) > 0 {
			c.Cells[position].Color = color
		} else {
			c.Cells[position].Color = CELL_COLOR_1
		}
	}
	c.Step <- struct{}{}
}

// Start starts infinite loop to run c.Run() per second
func (c *Controller) Start(ctx context.Context) {
	if c.IsStart {
		return
	}
	c.ctx, c.cancel = context.WithCancel(ctx)
	c.IsStart = true
	for {
		c.Run()
		c.Step <- struct{}{}
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
		c.cancel = nil
		c.IsStart = false
	}
}

// Next likes Run(), but if controller is running
// should stop current loop and run new loop after Run()
func (c *Controller) Next() {
	if c.IsStart {
		c.Stop()
		c.Run()
		c.Start(context.Background())
	} else {
		c.Run()
		c.Step <- struct{}{}
	}
}

// Reset is reset cells to default, and version to 0
func (c *Controller) Reset() {
	if c.IsStart {
		c.Stop()
	}

	c.Cells = initCells(c.Size)
	c.Version = 0
	c.IsStart = false
	c.Step <- struct{}{}
}

// NewController create a controller which initialize size*size cells
// and setup cell's neighbors.
func NewController(size int) *Controller {
	cells := initCells(size)
	step := make(chan struct{})
	return &Controller{Cells: cells, Size: size, Step: step}
}
// Init create controller with default shape
// size is the cells count (equal size*size)
func Init(size int) {
	CurrentController = NewController(size)
	SetToPatterns(CELL_PATTERN_GLIDER, size)
}

// initCells create a new cells (all die)
func initCells(size int) []*Cell {
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
	return cells
}

// SetToPatterns accepts size to calculate center position
// and put predefine special shape in the center.
// TODO: refactoring to accept position, but need to validate over edge or not
func SetToPatterns(pattern CellPattern, size int) {
	if CurrentController.IsStart {
		CurrentController.Stop()
	}
	selectedCell := make(map[int]bool)
	center := (size/2-1)*size + (size/2 - 1)
	switch pattern {
	case CELL_PATTERN_GLIDER:
		//	X O X
		//	X X O
		//	O O O
		selectedCell[center-size] = true
		selectedCell[center+1] = true
		selectedCell[center+size-1] = true
		selectedCell[center+size] = true
		selectedCell[center+size+1] = true
	case CELL_PATTERN_BEACON:
		//	O O X X
		//	O O X X
		//	X X O O
		//  X X O O
		selectedCell[center-size*2-1] = true
		selectedCell[center-size*2] = true
		selectedCell[center-size-1] = true
		selectedCell[center-size] = true
		selectedCell[center+1] = true
		selectedCell[center+2] = true
		selectedCell[center+size+1] = true
		selectedCell[center+size+2] = true

	case CELL_PATTERN_BOAT:
		//	O O X
		//	O X O
		//	X O X
		selectedCell[center-size-1] = true
		selectedCell[center-size] = true
		selectedCell[center-1] = true
		selectedCell[center+1] = true
		selectedCell[center+size] = true
	default:

	}
	for i, v := range CurrentController.Cells {
		if selectedCell[i] == true {
			v.Status = CELL_LIVE
			v.Color = CELL_COLOR_1
		} else {
			v.Status = CELL_DIE
		}
	}
}
