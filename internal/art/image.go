package art

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

type CardImage struct {
	Filename     string
	Width        int
	Height       int
	Colors       []RGB
	AverageColor RGB
}

func MakeImage(filename string) *CardImage {
	file := loadImageFile(filename)
	img := decodeImage(file)
	width := img.Bounds().Max.X - 1
	height := img.Bounds().Max.Y - 1
	colors := getColors(img)
	average := getAverageColor(colors)

	return &CardImage{
		Filename:     filename,
		Width:        width,
		Height:       height,
		Colors:       colors,
		AverageColor: average,
	}
}

func loadImageFile(filename string) *os.File {
	img, err := os.Open(filename)

	if err != nil {
		log.Fatal("Can't open jpeg: ", err)
	}
	return img
}

func decodeImage(file *os.File) image.Image {
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal("Unable to decode jpeg")
	}
	defer file.Close()
	return img
}

func getColors(img image.Image) []RGB {
	rgb := make([]RGB, 0)
	width := img.Bounds().Max.X - 1
	height := img.Bounds().Max.Y - 1

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			pixel := color.NRGBAModel.Convert(img.At(i, j))
			red, green, blue, _ := pixel.RGBA()
			color := RGB{
				Red:   uint8(red),
				Green: uint8(green),
				Blue:  uint8(blue),
			}
			rgb = append(rgb, color)
		}
	}
	return rgb
}

func getAverageColor(colors []RGB) RGB {
	redSum := 0
	greenSum := 0
	blueSum := 0
	length := len(colors)

	for _, v := range colors {
		redSum += int(v.Red)
		greenSum += int(v.Green)
		blueSum += int(v.Blue)
	}

	redSum /= length
	greenSum /= length
	blueSum /= length

	return RGB{
		Red:   uint8(redSum),
		Green: uint8(greenSum),
		Blue:  uint8(blueSum),
	}
}

func DetermineMostCommonColor(colors []RGB) string {
	freq := make(map[string]int)
	for _, v := range colors {
		color := v.DetermineColor()
		if freq[color] == 0 {
			freq[color] = 1
		} else {
			freq[color]++
		}
	}

	name := ""
	max := 0

	for k, v := range freq {
		if max < v {
			max = v
			name = k
		}
	}

	return name
}
