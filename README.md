# Multiplayer Conway’s Game of Life

[Conway’s Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) is a famous simulation that demonstrates cellular automaton. It is modeled as a grid with 4 simple rules:

1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.


This is a backend service for game of life, include managing cells life cycle, providing RESTful APIs, and Websocket server.


## Solution

### Controller & Cell
```
type Cell struct {
	Status     CellStatus 
	nextStatus CellStatus
	neighbors  []*Cell
}
```

cell keeps pointer of neighbors, when enter the evaluate stage, it would be easy to calculate the number of living neighbors.

```
type Controller struct {
	Cells   []*Cell
	Size    int
	Version int
	ctx     context.Context
	cancel  context.CancelFunc
	IsStart bool
	Step    chan struct{}
}
```

controller responses for managing 2 stages of cell's lifecycle:

1. evaluate: evaluating cell will live or die. 
2. refresh: changing cell's status to live or die.


### Web API

RESTful api provide start, stop, next, reset, reverse and get cells methods to client.

### Websocket
we use websocket to notify client to get update of cell's status, thus, client could not setting timer to fetch result, and easy to share other clients reverse action. 

## How to test

```
go test -v ./...
```

## How to build 
```
go build .
./gameoflife -addr :8080 -size 20
// addr is running address:port
// size is the cell size
```

## How to deploy

