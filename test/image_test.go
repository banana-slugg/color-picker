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
