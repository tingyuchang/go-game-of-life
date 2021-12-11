package model

import (
	"testing"
)

func TestNewController(t *testing.T) {
	cell0 := &Cell{Coordinate: Coordinate{X: 0, Y: 0}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell1 := &Cell{Coordinate: Coordinate{X: 0, Y: 1}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell2 := &Cell{Coordinate: Coordinate{X: 0, Y: 2}, Status: CELL_DIE, neighbors: []*Cell{}}

	cell3 := &Cell{Coordinate: Coordinate{X: 1, Y: 0}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell4 := &Cell{Coordinate: Coordinate{X: 1, Y: 1}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell5 := &Cell{Coordinate: Coordinate{X: 1, Y: 2}, Status: CELL_DIE, neighbors: []*Cell{}}

	cell6 := &Cell{Coordinate: Coordinate{X: 2, Y: 0}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell7 := &Cell{Coordinate: Coordinate{X: 2, Y: 1}, Status: CELL_DIE, neighbors: []*Cell{}}
	cell8 := &Cell{Coordinate: Coordinate{X: 2, Y: 2}, Status: CELL_DIE, neighbors: []*Cell{}}

	cell0.SetNeighbors([]*Cell{cell1, cell3, cell4})
	cell1.SetNeighbors([]*Cell{cell0, cell2, cell3, cell4, cell5})
	cell2.SetNeighbors([]*Cell{cell1, cell4, cell5})
	cell3.SetNeighbors([]*Cell{cell0, cell1, cell4, cell6, cell7})
	cell4.SetNeighbors([]*Cell{cell0, cell1, cell2, cell3, cell5, cell6, cell7, cell8})
	cell5.SetNeighbors([]*Cell{cell1, cell2, cell4, cell7, cell8})
	cell6.SetNeighbors([]*Cell{cell3, cell4, cell7})
	cell7.SetNeighbors([]*Cell{cell3, cell4, cell5, cell6, cell8})
	cell8.SetNeighbors([]*Cell{cell4, cell5, cell7})

	want := &Controller{Size: 3, Version: 0, Cells: []*Cell{cell0, cell1, cell2, cell3, cell4, cell5, cell6, cell7, cell8}}

	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want *Controller
	}{
		{
			name: "test create 3*3 cells, cell neighbors count equal want's",
			args: args{
				size: 3,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewController(tt.args.size)

			for i, v := range got.Cells {
				validCellCount := 0
				for _, v2 := range v.neighbors {
					if v2 != nil {
						validCellCount++
					}
				}
				cell := tt.want.Cells[i]
				if len(cell.neighbors) != validCellCount {
					t.Errorf("Neighbors = %d, want %d", validCellCount, len(cell.neighbors))
				}
			}
		})
	}
}
