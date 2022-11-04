package main

import (
	"fmt"

	"github.com/crims1n/color-picker/internal/art"
	"github.com/crims1n/color-picker/internal/util"
)

func main() {
	colors := [5]string{"white", "blue", "black", "red", "green"}

	for _, color := range colors {
		path := "./assets/" + color
		files := util.FilesFromDir(path)
		avg := 0
		freq := 0
		resAvg := make(chan string)
		resFreq := make(chan string)
		for _, v := range files {
			go func(v string) {
				filename := fmt.Sprintf("./assets/%v/%v", color, v)
				img := art.MakeImage(filename)
				freq := art.DetermineMostCommonColor(img.Colors)

				resAvg <- img.AverageColor.DetermineColor()
				resFreq <- freq
			}(v)
		}

		for i := 0; i < 2000; i++ {
			select {
			case a := <-resAvg:
				if a == color {
					avg++
				}
			case f := <-resFreq:
				if f == color {
					freq++
				}
			}
		}

		percentAvg := float64(avg) / 10.0
		percentFreq := float64(freq) / 10.0
		fmt.Printf("%v:\t Average:%v%%\t Frequency:%v%%\n", color, percentAvg, percentFreq)
	}

}
