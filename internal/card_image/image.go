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

func DecodeImage(file *os.File) {
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal("Unable to decode jpeg")
	}
	defer file.Close()

	fmt.Printf("%v\n: ", img.Bounds())
	color := color.NRGBAModel.Convert(img.At(1, 1))
	fmt.Println(color)

}
