package art

import (
	"fmt"
	"math"

	"github.com/crims1n/color-picker/internal/util"
)

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func (c RGB) PrintValues() {
	fmt.Printf("%v, %v, %v\n", c.Red, c.Green, c.Blue)
}

func (c RGB) DetermineColor() string {
	thresh := 65
	r := c.Red
	g := c.Green
	b := c.Blue
	max := util.Max(r, g, b)
	avg := (r + g + b) / 3

	rg := math.Abs(float64(r - g))
	rb := math.Abs(float64(r - b))
	gb := math.Abs(float64(g - b))

	diff := util.Max(rg, rb, gb)

	if diff <= 25 {
		if max < uint8(thresh) || avg < uint8(thresh) {
			return "black"
		} else {
			return "white"
		}
	}

	if r >= util.Max(g, b) {
		return "red"
	}

	if g >= util.Max(r, b) {
		return "green"
	}

	if b >= util.Max(r, g) {
		return "blue"
	}

	return "black"
}
