# Multiplayer Conway’s Game of Life

[Conway’s Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life) is a famous simulation that demonstrates cellular automaton. It is modeled as a grid with 4 simple rules:

1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

For this repository

1. Implement the Game of Life browser frontend. You can use any representation such as `canvas`, simple DOM manipulation or even `table` cells. The game should tick automatically at a predefined interval, at say, 1 step per second.
2. The browser connects to a server, which allows multiple browsers to share the same, synchronized world view. Unless otherwise specified, the server may be written in Ruby, Node.js, or any other technology supported by Heroku. You may use any framework, e.g. Ruby on Rails, Hapi, Phoenix or just plain listening on a socket.
3. Each client is assigned a random color on initialization. From the browser, clicking on any grid will create a live cell on that grid with the client’s color. This change should be synchronized across all connected clients. (You can use any mechanism to achieve this, such as polling, comet, or WebSocket.)
4. When a dead cell revives by rule #4 “Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.”, it will be given a color that is the average of its neighbors (that revive it).
5. To make the evolution more interesting, include a toolbar that places some predefined patterns at random places with the player’s color, such as those found at here https://en.wikipedia.org/wiki/Conway’s_Game_of_Life#Examples_of_patterns (not necessary to implement all, just 3 - 4 is fine).


This is a backend service for game of life, include managing cells life cycle, providing RESTful APIs, and Websocket server.

### Demo
[demo site](https://afternoon-fjord-92266.herokuapp.com/)
[frontend repo](https://github.com/tingyuchang/vue-game-of-life)

![Imgur](https://imgur.com/SwC0P2u.gif)


## Solution

### Controller & Cell

```
type Cell struct {
	Status     CellStatus 
	nextStatus CellStatus
	neighbors  []*Cell
}
```

Cell keeps pointer of neighbors, when enter the evaluate stage, it would be easy to calculate the number of living neighbors.
For #4 requirement, it's also easy determine average color from neighbors.

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

Actually, controller handle all behavior in this solution, like:

1. start auto-step
2. stop auto-stop
3. jump into next step
4. reset all cells
5. revive or make cell die


### Web API

RESTful api provide start, stop, next, reset, reverse and get cells methods to client to control controller.

### Websocket
we use websocket to notify client to get update of cell's status, thus, client could not setting timer to fetch result, and easy to share other clients reverse action. 

## How to test

We provide a stdout method to easy check cells status 
`model.Controller.Show()`

For unit testing:
```
go test -v ./...
```

## How to build 

```
go build .
./gameoflife -addr 8080 -size 20
// addr is running address
// size is the cell size
```

## How to deploy

### heroku

We recommend using docker to deploy this app, it could be easy to transfer in different platforms.

```
// heroku cli
heroku create
heroku container:push web 
// web is the name you can change it to anything you want
heroku container:release web
heroku ps:scale web=1
heroku open
```

## Discussion

**Q: If the client disconnects and reconnects some time in the future, do they need to keep the same color?**

Current solution reassign random color when client reload the page.
If we want to keep the same color, there are 2 ways to achieve:

1. server keeps client's identify info, when the client ask a random color, we can make sure give a same color.
2. client keeps color in local storage

**Q: every single update sends all cells info, can we just get updated cells?**

it seems good for transfer efficiency, but current process is push base, server is hard to know client's status (or version), if change to pull base, client would `pull` much more requests than push base and server need to handle version diff for each request, i don't think this way is more efficiency. BUT maybe i'm wrong, welcome to discuss.

**How to scale?**

We can't horizontal scale currently solution, only vertical scaling, because we keeps data in memory, thus if multiple instances exist, each server has own cells, that is not we want. To achieve horizontal scaling, these are 2 ways:

1. separating package model from api & websocket to a single service, when users grow, connections use to be a bottleneck, api & websocket could scale horizontal.
2. Do not keep data in RAM, use Redis instead.

## Reference 

- https://github.com/gorilla/websocket
  