package main

import (
	"gameoflife/model"
)

var size = 10

func main() {
	c := model.NewController(size)
	c.Show()
	c.Reverse(0)
	c.Reverse(1)
	c.Reverse(10)
	c.Reverse(11)
	c.Reverse(22)
	c.Reverse(23)
	c.Reverse(32)
	c.Reverse(33)
	c.Show()
	c.Run()
	c.Show()
	c.Run()
	c.Show()
	c.Run()
	c.Show()
	c.Run()
	c.Show()
	c.Run()
	c.Show()
}
