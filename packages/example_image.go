package packages

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

/* Goal: Define your own Image type, implement the necessary methods, and call pic.ShowImage. */
// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
// ColorModel should return color.RGBAModel.
// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

type Image struct{}

// Draws using RGBA (4 bytes) for each pixel
func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Defines size of the image
// 0, 0 - top-left
// 256, 256 - bottom-right
func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// Assigns color for each pixel
func (Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func ExampleImage() {
	m := Image{}
	pic.ShowImage(m)
}
