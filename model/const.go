package model

import "fmt"

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

var CurrentController *Controller

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

const (
	CELL_COLOR_1  = "#000000"
	CELL_COLOR_2  = "#454545"
	CELL_COLOR_3  = "#999999"
	CELL_COLOR_4  = "#FF0000"
	CELL_COLOR_5  = "#00FF00"
	CELL_COLOR_6  = "#0000FF"
	CELL_COLOR_7  = "#800000"
	CELL_COLOR_8  = "#808000"
	CELL_COLOR_9  = "#008000"
	CELL_COLOR_10 = "#00FFFF"
	CELL_COLOR_11 = "#000080"
	CELL_COLOR_12 = "#FF00FF"
	CELL_COLOR_13 = "#800080"
	CELL_COLOR_14 = "#DFFF00"
	CELL_COLOR_15 = "#9FE2BF"
)
