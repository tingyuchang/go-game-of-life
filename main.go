package main

import (
	"context"
	"gameoflife/model"
	"net/http"
	"time"
)

var size = 10

func main() {
	c := model.GetStartWithGlider(size)
	go c.Start(context.Background())

	go func() {
		time.Sleep(10 * time.Second)
		c.Stop()
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
