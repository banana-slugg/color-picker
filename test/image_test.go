package image

import (
	"fmt"
	"testing"

	"github.com/crims1n/color-picker/internal/card_image"
)

func TestImageDecode(t *testing.T) {
	img := card_image.LoadImageFile("./testdata/is-black-whitecorner.jpg")
	fmt.Println(card_image.DecodeImage(img))
}
