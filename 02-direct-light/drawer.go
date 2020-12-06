package main

import (
	"image"
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
)

func gcDrawCanvas(gc *draw2dimg.GraphicContext, c Canvas) {
	gcDrawRectangle(gc, c.Outline.Min, c.Outline.Max, NoLightColor)

	for _, b := range c.Beams {
		gcDrawBeamSource(gc, b)
	}

	for _, w := range c.Shapes {
		gcDrawShape(gc, w)
	}
	gc.Close()
	gc.FillStroke()
}

func gcDrawBeamSource(gc *draw2dimg.GraphicContext, b Beam) {
	p := b.source
	gcDrawRectangle(gc, image.Point{p.X - 4, p.Y - 4}, image.Point{p.X + 4, p.Y + 4}, FlashLightColor)
	// draw light source internal box
	gcDrawRectangle(gc, image.Point{p.X - 1, p.Y - 1}, image.Point{p.X + 1, p.Y + 1}, FullLightColor)
	// draw light beam
	gcDrawBeam(gc, b)
}

func gcDrawBeam(gc *draw2dimg.GraphicContext, b Beam) {
	gcDrawLine(gc, b.source, b.Ends(b.canvas.Shapes), FullLightColor)
}

func gcDrawShape(gc *draw2dimg.GraphicContext, w image.Rectangle) {
	gcDrawLine(gc, w.Min, w.Max, WallColor)
}

func gcDrawLine(gc *draw2dimg.GraphicContext, start image.Point, end image.Point, c color.Color) {
	gc.SetStrokeColor(c)
	gc.SetLineWidth(2)

	gc.BeginPath() // Initialize a new path
	gc.MoveTo(float64(start.X), float64(start.Y))
	gc.LineTo(float64(end.X), float64(end.Y))
	gc.Close()

	gc.Stroke()
}

func gcDrawRectangle(gc *draw2dimg.GraphicContext, start image.Point, end image.Point, c color.Color) {
	draw2dkit.Rectangle(gc, (float64)(start.X), (float64)(start.Y), (float64)(end.X), (float64)(end.Y))
	gc.SetStrokeColor(c)
	gc.SetLineWidth(2)
	gc.Stroke()
}
