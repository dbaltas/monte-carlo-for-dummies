package main

import (
	"image"
	"math"
	"testing"
)

var delta = 0.001

var testAngles = []struct {
	p1       image.Point
	p2       image.Point
	expAngle float64
}{
	{image.Point{0, 0}, image.Point{0, 0}, 0.},
	// top left
	{image.Point{0, 0}, image.Point{1, 0}, 0.},
	{image.Point{0, 0}, image.Point{1, 1}, 0.785}, // 45. deg
	// top right
	{image.Point{0, 0}, image.Point{0, 1}, 1.571},  // 90. deg
	{image.Point{0, 0}, image.Point{-1, 1}, 2.356}, // 135. deg
	// bottom right
	{image.Point{0, 0}, image.Point{-1, 0}, 3.142},   // 180. deg
	{image.Point{0, 0}, image.Point{-1, -1}, -2.356}, // -135. deg
	// bottom left
	{image.Point{0, 0}, image.Point{0, -1}, -1.571}, // -90. deg
	{image.Point{0, 0}, image.Point{1, -1}, -0.785}, // -45. deg
}

func TestAngle(t *testing.T) {
	for _, e := range testAngles {
		angle := Angle(e.p1, e.p2)
		if e.expAngle-angle > delta {
			t.Errorf("Exp angle %f got %f for %v to %v", e.expAngle, angle,
				e.p1, e.p2)
		}
	}
}

var testVerticalThroughCenter = []struct {
	p1 image.Point  // line start
	p2 image.Point  // line end
	pi *image.Point // exp intersection point
}{
	{image.Point{0, -10}, image.Point{0, 10}, &image.Point{0, 0}},
	{image.Point{-5, 5}, image.Point{5, -5}, &image.Point{0, 0}},
	{image.Point{-100, -10}, image.Point{200, 10}, &image.Point{49, 0}},
	{image.Point{-100, -10}, image.Point{100, -10}, nil},
	{image.Point{1, 1}, image.Point{2, 2}, &image.Point{0, 0}}, // line extension intersects
}

func TestLineIntersects(t *testing.T) {
	p1 := image.Point{-10000, 0}
	p2 := image.Point{10000, 0}
	l := Line{p1, p2}
	for _, e := range testVerticalThroughCenter {
		l1 := Line{e.p1, e.p2}
		pi := l.Intersection(l1)

		if !cmpEqual(e.pi, pi) {
			t.Errorf("Exp intersection point: %v got %v for lines %v,  %v",
				e.pi, pi, l, l1)
		}
	}
}

var testAngleConversion = []struct {
	rad float64 // line start
	deg float64 // line end
}{
	{0, 0},
	{math.Pi / 2, 90},
	{math.Pi, 180},
	{-math.Pi, -180},
	{-math.Pi / 2, -90},
}

func TestAngleMeasurements(t *testing.T) {
	// DegToRad
	for _, e := range testAngleConversion {
		if math.Abs(e.rad-DegToRad(e.deg)) > delta {
			t.Errorf("Exp DegToRad: %v got %v for deg %v and delta %v",
				e.rad, DegToRad(e.deg), e.deg, delta)
		}
	}
	// RadToDeg
	for _, e := range testAngleConversion {
		if math.Abs(e.deg-RadToDeg(e.rad)) > delta {
			t.Errorf("Exp RadToDeg: %v got %v for rad %v and delta %v",
				e.deg, RadToDeg(e.rad), e.rad, delta)
		}
	}
}

var testAngleToPoint = []struct {
	rad float64 // angle in radius
	p   image.Point
}{
	{0, image.Point{1000, 0}},
	{math.Pi / 2, image.Point{0, 1000}},
	{math.Pi, image.Point{-1000, 0}},
	{-math.Pi / 2, image.Point{0, -1000}},
}

func TestAngleToPoint(t *testing.T) {
	distance := 1000.
	// test distances
	for _, e := range testAngleToPoint {
		p := image.Point{0, 0}
		// p2 := image.Point{
		// 	X: int(distance*math.Cos(e.rad)) + p.X,
		// 	Y: int(distance*math.Sin(e.rad)) + p.Y,
		// }
		p2 := PointInDirection(p, e.rad, distance)
		if !cmpEqual(&e.p, &p2) {
			t.Errorf("exp %v got %v for angle %v and distance %v",
				e.p, p2, RadToDeg(e.rad), distance)
		}
	}
}

// compare structs, avoiding using reflection.deepEqual
// consider adding an external dependency like testify or go-cmp
func cmpEqual(p1 *image.Point, p2 *image.Point) bool {
	if p1 == nil && p2 == nil {
		// no intersection point
		return true
	}
	if p1 == nil || p2 == nil {
		return false
	}
	if p1.X == p2.X && p1.Y == p2.Y {
		return true
	}
	return false
}
