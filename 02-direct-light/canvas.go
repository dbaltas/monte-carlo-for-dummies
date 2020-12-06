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
	outline := image.Rect(0, 0, 1600, 900)
	c := Canvas{
		Outline: outline,
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		b := Beam{
			source:    image.Point{r.Intn(outline.Max.X), r.Intn(outline.Max.Y)},
			angle:     float64(r.Intn(1000)) / 6.2830,
			intensity: 100,
			canvas:    &c,
		}
		c.Beams = append(c.Beams, b)
	}

	for i := 0; i < 10; i++ {
		randX1 := r.Intn(outline.Max.X)
		randX2 := r.Intn(outline.Max.X)
		randY1 := r.Intn(outline.Max.Y)
		randY2 := r.Intn(outline.Max.Y)
		w := image.Rect(randX1, randY1, randX2, randY2)
		c.Shapes = append(c.Shapes, w)
	}

	return c
}
