package main

import (
	"image"
	"math"
)

// BeamAngleSpread The beam is more of a triangle instead of a straight line
var BeamAngleSpread = 0.03 * math.Pi

// Beam A flow of photons
type Beam struct {
	source    image.Point
	angle     float64
	intensity int
	canvas    *Canvas
}

// Targets True if the direction of the Beam targets the point
func (b Beam) Targets(p image.Point) bool {

	diff := math.Abs(b.angle - Angle(b.source, p))
	if diff < BeamAngleSpread {
		return true
	}
	if (2*math.Pi - diff) < BeamAngleSpread {
		return true
	}

	return false
}

// EndPoint for rendering purposes
func (b Beam) EndPoint() image.Point {
	distance := 100000.
	return PointInDirection(b.source, b.angle, distance)
}

// Ends for rendering purposes
func (b Beam) Ends(walls []image.Rectangle) image.Point {
	shorterEnd := b.EndPoint()
	l := Line{b.source, shorterEnd}
	for _, w := range walls {
		// TODO: hack gets the diagonal of the rect instead of all lines
		l1 := Line{w.Min, w.Max}
		p := l.Intersection(l1)
		if p == nil {
			continue
		}
		if Distance(b.source, *p) < Distance(b.source, shorterEnd) {
			shorterEnd = *p
		}
	}
	return shorterEnd
}
