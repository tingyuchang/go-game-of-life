package main

import (
	"fmt"
	"gameoflife/model"
)

var size = 10

func main() {
	c := model.NewController(size)
	fmt.Println(c)
}
