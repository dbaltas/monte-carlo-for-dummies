package main

import (
	"image"
	// "math"
	"testing"
)

func TestProperties(t *testing.T) {
	beams := []Beam{}
	shapes := []image.Rectangle{}
	r := image.Rect(0, 0, 125, 32)
	_ = Canvas{
		Outline: r,
		Beams: beams,
		Shapes: shapes,
	}
}
