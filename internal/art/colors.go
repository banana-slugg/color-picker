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
	// find the max distance between the RGB values
	r := c.Red
	g := c.Green
	b := c.Blue
	maxVal := util.Max(r, g, b)

	rg := math.Abs(float64(r - g))
	rb := math.Abs(float64(r - b))
	gb := math.Abs(float64(g - b))

	maxDiff := util.Max(rg, rb, gb)

	if maxDiff <= 30 {
		if maxVal < 70 {
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

	return "blue"

}
