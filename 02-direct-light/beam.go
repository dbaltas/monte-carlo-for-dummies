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
	distance := 1000.
	return PointInDirection(b.source, b.angle, distance)
}
