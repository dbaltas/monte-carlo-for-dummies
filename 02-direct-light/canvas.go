package main

import (
	"image"
	"math/rand"
	"time"
)

// Canvas A board with outline beams and shapes
type Canvas struct {
	Outline image.Rectangle
	Beams   []Beam
	Shapes  []image.Rectangle
}

// RandomCanvas render random beams and shapes
func RandomCanvas() Canvas {
	outline := image.Rect(0, 0, 2500, 640)
	c := Canvas{
		Outline: outline,
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 200; i++ {
		b := Beam{
			source:    image.Point{r.Intn(outline.Max.X), r.Intn(outline.Max.Y)},
			angle:     float64(r.Intn(1000)) / 6.2830,
			intensity: 100,
		}
		c.Beams = append(c.Beams, b)
	}

	for i := 0; i < 5; i++ {
		randX := r.Intn(outline.Max.X)
		w := image.Rect(randX, r.Intn(outline.Max.Y), randX, r.Intn(outline.Max.Y))
		c.Shapes = append(c.Shapes, w)
	}

	return c
}
