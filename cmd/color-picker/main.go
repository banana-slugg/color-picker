package main

import (
	"fmt"

	"github.com/crims1n/color-picker/internal/art"
	"github.com/crims1n/color-picker/internal/util"
)

func main() {
	colors := [5]string{"white", "blue", "black", "red", "green"}
	for group := 0; group < 10; group++ {
		groupFrequencyAcc := 0
		groupAverageAcc := 0
		averageChan := make(chan string)
		frequencyChan := make(chan string)
		for _, color := range colors {
			path := fmt.Sprintf("./assets/group-%v/validation/%v", group, color)
			files := util.FilesFromDir(path)
			for _, f := range files {
				file := fmt.Sprintf("%v/%v", path, f)
				go runner(file, color, averageChan, frequencyChan)
			}

			for i := 0; i < 200; i++ {
				select {
				case a := <-averageChan:
					if a == color {
						groupAverageAcc++
					}
				case f := <-frequencyChan:
					if f == color {
						groupFrequencyAcc++
					}
				}
			}

		}
		percentAvg := float64(groupAverageAcc) / 5
		percentFreq := float64(groupFrequencyAcc) / 5

		fmt.Printf("Group %v:\t Average:%v%%\t Frequency:%v%%\n", group, percentAvg, percentFreq)
	}

}

func runner(filename string, color string, resAvg chan string, resFreq chan string) {
	img := art.MakeImage(filename)
	freq := art.DetermineMostCommonColor(img.Colors)

	resAvg <- img.AverageColor.DetermineColor()
	resFreq <- freq
}
