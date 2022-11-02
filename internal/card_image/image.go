package card_image

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

type CardImage struct {
	Filename string
	Width    int
	Height   int

	colors RGB
}

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func LoadImageFile(filename string) *os.File {
	img, err := os.Open(filename)

	if err != nil {
		log.Fatal("Can't open jpeg")
	}

	return img
}

func DecodeImage(file *os.File) string {
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal("Unable to decode jpeg")
	}
	defer file.Close()

	color := color.NRGBAModel.Convert(img.At(625, 456))

	width := img.Bounds().Max.X - 1 //img.Bounds() is zero indexed, and the max is not inclusive, so minus 1
	height := img.Bounds().Max.Y - 1
	return fmt.Sprintf("width: %v\nheight: %v\nRGB: %v", width, height, color)

}
