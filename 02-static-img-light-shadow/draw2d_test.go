package main

import (
	"image"
	"image/color"
	"testing"

	"github.com/llgcode/draw2d/draw2dimg"
)

func TestVendorDraw2d(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 5))

	// sanity check, all black!
	assertPixel(t, img, image.Point{0, 0}, color.RGBA{0x00, 0x00, 0x00, 0x00})
	assertPixel(t, img, image.Point{20, 3}, color.RGBA{0x00, 0x00, 0x00, 0x00})

	gc := draw2dimg.NewGraphicContext(img)

	gc.SetFillColor(color.RGBA{0xaa, 0xaa, 0xaa, 0xff})
	gc.SetStrokeColor(color.RGBA{0xff, 0xbb, 0xbb, 0xff})
	gc.SetLineWidth(1)
	gc.BeginPath()   // Initialize a new path
	gc.MoveTo(10, 3) // Move to a position to start the new path
	gc.LineTo(90, 3)
	gc.Close()
	gc.FillStroke()

	// 0,0 stays black!
	assertPixel(t, img, image.Point{0, 0}, color.RGBA{0x00, 0x00, 0x00, 0x00})
	// 20,3 changes color!
	assertPixel(t, img, image.Point{20, 3}, color.RGBA{0xff, 0xbb, 0xbb, 0xff})
}

func assertPixel(t *testing.T, img image.Image, p image.Point, c color.Color) {
	if c != img.At(p.X, p.Y) {
		t.Errorf("exp: %t, got %t at %v", c, img.At(p.X, p.Y), p)
	}
}
