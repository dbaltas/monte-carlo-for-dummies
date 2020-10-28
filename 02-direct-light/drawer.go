package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
)

func gcDrawLine(gc *draw2dimg.GraphicContext, start image.Point, end image.Point, c color.Color) {
	// running 4 times! since this method call produces different results when called multiple times
	for i := 0; i < 4; i++ {
		gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0x00})
		gc.SetStrokeColor(c)
		gc.SetLineWidth(1)

		gc.BeginPath() // Initialize a new path
		gc.MoveTo(float64(start.X), float64(start.Y))
		gc.LineTo(float64(end.X), float64(end.Y))

		gc.Close()
		gc.Stroke()
	}
}

func gcDrawRectangle(gc *draw2dimg.GraphicContext, start image.Point, end image.Point, c color.Color) {
	// running 4 times! since this method call produces different results when called multiple times
	for i := 0; i < 4; i++ {
		gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0x00})
		gc.SetStrokeColor(c)
		gc.SetLineWidth(1)

		gc.BeginPath() // Initialize a new path
		gc.MoveTo(float64(start.X), float64(start.Y))
		gc.LineTo(float64(end.X), float64(start.Y))
		gc.LineTo(float64(end.X), float64(end.Y))
		gc.LineTo(float64(start.X), float64(end.Y))
		gc.LineTo(float64(start.X), float64(start.Y))

		gc.Close()
		gc.Stroke()
	}
}
