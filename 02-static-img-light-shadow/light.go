package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

// Pixel in 2 dimensions
type Pixel struct {
	x int
	y int
}

func main() {
	err := createCanvasPngFile("canvas.png")
	if err != nil {
		log.Fatal(err)
	}

	renderASCIIGrayScale("canvas.png")
}

func createCanvasPngFile(outFile string) error {
	myImg := image.NewRGBA(image.Rect(0, 0, 125, 32))

	for i := 0; i < len(myImg.Pix); i++ {
		if i%4 == 0 {
			myImg.Pix[i+3] = 255
		}
	}

	flashLightCenterOffset := 8060
	i := flashLightCenterOffset

	myImg.Pix[i] = 255
	myImg.Pix[i+1] = 0
	myImg.Pix[i+2] = 0

	out, err := os.Create(outFile)

	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, myImg)

	return err
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

	// levels := []string{" ", "░", "▒", "▓", "█"}
	levels := []string{"░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			// level := c.Y / 51 // 51 * 5 = 255
			level := c.Y / 61
			if level >= 4 {
				level = 3
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
