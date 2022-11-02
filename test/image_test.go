package image

import (
	"testing"

	"github.com/crims1n/color-picker/internal/card_image"
)

func TestImageDecode(t *testing.T) {
	img := card_image.LoadImageFile("./testdata/is-blue.jpg")
	card_image.DecodeImage(img)
}
