package util

import (
	"fmt"
	"image/color"
	"strconv"
)

// GetAverageColor
// #000000
func GetAverageColor(colors []string) string {
	colorNum := 0
	rgbColors := make([]color.RGBA, len(colors))
	for i, v := range colors {
		if v != "" {
			rgbColors[i], _ = parseHexColor(v)
			colorNum++
		}
	}
	var r float64
	var g float64
	var b float64

	for _, v := range rgbColors {
		r += float64(v.R)
		g += float64(v.G)
		b += float64(v.B)
	}
	r = r / float64(colorNum)
	g = g / float64(colorNum)
	b = b / float64(colorNum)

	return fmt.Sprintf("#%v%v%v", floatToHex(r), floatToHex(g), floatToHex(b))
}

func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

func floatToHex(f float64) string {
	result := strconv.FormatInt(int64(f), 16)
	if len(result) == 1 {
		result = "0" + result
	}
	return result
}
