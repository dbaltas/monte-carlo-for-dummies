package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
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

func debug(img image.Image) {
	levels := []string{"░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			// level := c.Y / 61
			if level >= 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("boo\n")
	}
	fmt.Print(img)
	log.Printf("%v", img.Bounds())
	log.Printf("%v", img.Bounds().Min)
	log.Printf("%v", img.Bounds().Max)
	log.Printf("y:%d x:$%d", img.Bounds().Max.Y, img.Bounds().Max.X)
}
