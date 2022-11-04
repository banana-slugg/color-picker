package art_test

import (
	"fmt"
	"testing"

	"github.com/crims1n/color-picker/internal/art"
)

func TestBlackImage(t *testing.T) {
	img := art.MakeImage("./testdata/is-black.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
}

func TestWhiteImage(t *testing.T) {
	img := art.MakeImage("./testdata/is-white.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
	fmt.Println(art.DetermineMostCommonColor(img.Colors))
}

func TestGreenImage(t *testing.T) {
	img := art.MakeImage("./testdata/is-green.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
	fmt.Println(art.DetermineMostCommonColor(img.Colors))
}

func TestWhiteCard(t *testing.T) {
	img := art.MakeImage("./testdata/should-be-white.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
	fmt.Println(art.DetermineMostCommonColor(img.Colors))
}

func TestBlueCard(t *testing.T) {
	img := art.MakeImage("./testdata/should-be-blue.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
	fmt.Println(art.DetermineMostCommonColor(img.Colors))
}

func TestBlackCard(t *testing.T) {
	img := art.MakeImage("./testdata/should-be-black.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
	art.DetermineMostCommonColor(img.Colors)
}

func TestRedCard(t *testing.T) {
	img := art.MakeImage("./testdata/should-be-red.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
}

func TestGreenCard(t *testing.T) {
	img := art.MakeImage("./testdata/should-be-green.jpg")
	img.AverageColor.PrintValues()
	fmt.Println(img.AverageColor.DetermineColor())
}

// func TestAverageWhite(t *testing.T) {
// 	avgR := 0
// 	avgG := 0
// 	avgB := 0

// 	path := "../assets/white"
// 	files := util.FilesFromDir(path)
// 	for _, v := range files {
// 		filename := fmt.Sprintf("../assets/white/%v", v)
// 		img := art.MakeImage(filename)
// 		avgR += int(img.AverageColor.Red)
// 		avgG += int(img.AverageColor.Green)
// 		avgB += int(img.AverageColor.Blue)
// 	}

// 	avgR /= 1000.0
// 	avgG /= 1000.0
// 	avgB /= 1000.0

// 	fmt.Println(avgR, avgG, avgB)
// }

// func TestAverageBlack(t *testing.T) {
// 	avgR := 0
// 	avgG := 0
// 	avgB := 0

// 	path := "../assets/black"
// 	files := util.FilesFromDir(path)
// 	for _, v := range files {
// 		filename := fmt.Sprintf("../assets/black/%v", v)
// 		img := art.MakeImage(filename)
// 		avgR += int(img.AverageColor.Red)
// 		avgG += int(img.AverageColor.Green)
// 		avgB += int(img.AverageColor.Blue)
// 	}

// 	avgR /= 1000.0
// 	avgG /= 1000.0
// 	avgB /= 1000.0

// 	fmt.Println(avgR, avgG, avgB)
// }
