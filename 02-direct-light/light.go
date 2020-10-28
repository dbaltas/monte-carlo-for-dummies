package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	createCanvas("canvas.png")

	renderASCIIGrayScale("canvas.png")
}

func createCanvas(outFile string) error {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 125, 32))

	for i := 0; i < len(dest.Pix); i++ {
		if i%4 == 0 {
			dest.Pix[i] = CanvasColor.R
			dest.Pix[i+1] = CanvasColor.G
			dest.Pix[i+2] = CanvasColor.B
			dest.Pix[i+3] = CanvasColor.A
		}
	}

	gc := draw2dimg.NewGraphicContext(dest)
	// draw light source external box
	gcDrawRectangle(gc, image.Point{21, 11}, image.Point{29, 19}, FlashLightColor)
	// draw light source internal box
	gcDrawRectangle(gc, image.Point{24, 14}, image.Point{26, 16}, FullLightColor)
	// draw light beam
	gcDrawLine(gc, image.Point{31, 15}, image.Point{125, 15}, FullLightColor)
	// draw wall
	gcDrawRectangle(gc, image.Point{0, 0}, image.Point{125, 32}, NoLightColor)

	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile(outFile, dest)

	return nil
}

func renderASCIIGrayScale(file string) {
	// This example uses png.Decode which can only decode PNG images.
	canvasFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer canvasFile.Close()

	// Consider using the general image.Decode as it can sniff and decode any registered image format.
	img, err := png.Decode(canvasFile)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(img)

	levels := []string{" ", "░", "▒", "▓", "█"}
	// levels := []string{"░", "▒", "▓", "█"}

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
		fmt.Print("\n")
	}

	log.Printf("%v", img.Bounds())
	log.Printf("%v", img.Bounds().Min)
	log.Printf("%v", img.Bounds().Max)
	log.Printf("y:%d x:$%d", img.Bounds().Max.Y, img.Bounds().Max.X)
}
