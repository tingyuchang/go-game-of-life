package model

import "testing"

func TestCell_CheckLife(t *testing.T) {
	type fields struct {
		Coordinate Coordinate
		Status     CellStatus
		nextStatus CellStatus
		neighbors  []*Cell
	}
	case1 := fields{Status: CELL_LIVE, neighbors: []*Cell{{Status: CELL_LIVE}}}
	case2 := fields{Status: CELL_LIVE, neighbors: []*Cell{{Status: CELL_LIVE}, {Status: CELL_LIVE}}}
	case3 := fields{Status: CELL_LIVE, neighbors: []*Cell{{Status: CELL_LIVE}, {Status: CELL_LIVE}, {Status: CELL_LIVE}, {Status: CELL_LIVE}}}
	case4 := fields{Status: CELL_DIE, neighbors: []*Cell{{Status: CELL_LIVE}, {Status: CELL_LIVE}, {Status: CELL_LIVE}}}

	tests := []struct {
		name   string
		fields fields
		want   CellStatus
	}{
		{
			name:   "cell live, neighbor's live under 2, cell die",
			fields: case1,
			want:   CELL_DIE,
		},
		{
			name:   "cell live, neighbor's live equal 2 or 3, cell live",
			fields: case2,
			want:   CELL_LIVE,
		},
		{
			name:   "cell live, neighbor's live over 3, cell die",
			fields: case3,
			want:   CELL_DIE,
		},
		{
			name:   "cell die, neighbor's live equal 3, cell live",
			fields: case4,
			want:   CELL_LIVE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cell{
				Coordinate: tt.fields.Coordinate,
				Status:     tt.fields.Status,
				nextStatus: tt.fields.nextStatus,
				neighbors:  tt.fields.neighbors,
			}

			c.CheckLife()
			if c.nextStatus != tt.want {
				t.Errorf("Cell CheckLife = %v, want %v", c.nextStatus, tt.want)
			}
		})
	}
}
