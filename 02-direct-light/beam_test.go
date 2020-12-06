package main

import (
	"image"
	"math"
	"testing"
)

var testBeamTargets = []struct {
	beamAngle float64
	p         image.Point
	exp       bool
}{
	{0., image.Point{100, 0}, true},               // --> to the right
	{0., image.Point{100, 20}, false},             // --> to the right
	{math.Pi, image.Point{-120, 0}, true},         // --> to the left
	{math.Pi - 0.02, image.Point{-120, -1}, true}, // --> to the left, edge case +181 - 179
	{math.Pi / 2, image.Point{0, 55}, true},       // ^ upwards
	{-math.Pi / 2, image.Point{0, -55}, true},     // downwards
	{0.644 - 0.15, image.Point{4, 3}, false},      // point with angle 36.9
	{0.644 - 0.02, image.Point{4, 3}, true},       // point with angle 36.9
	{0.644, image.Point{4, 3}, true},              // point with angle 36.9
	{0.644 - 0.02, image.Point{4, 3}, true},       // point with angle 36.9
	{0.644 - 0.15, image.Point{4, 3}, false},      // point with angle 36.9
}

func TestBeamTargetsPoint(t *testing.T) {
	for _, e := range testBeamTargets {
		b := Beam{image.Point{0, 0}, e.beamAngle, 33, nil}
		res := b.Targets(e.p)
		if res != e.exp {
			t.Logf("%v", Angle(b.source, e.p))
			t.Errorf("Exp %v got %v On %v targeting %v", e.exp, res, b, e.p)
		}
	}
}

var testBeamEnd = []struct {
	beamStart image.Point
	walls     []image.Rectangle
	exp       image.Point
}{
	{
		image.Point{0, 25},
		[]image.Rectangle{
			image.Rect(40, 0, 40, 100),
			image.Rect(20, 0, 20, 100),
			image.Rect(60, 0, 60, 100),
		},
		image.Point{20, 25},
	},
}

func TestBeamFirstWall(t *testing.T) {
	for _, e := range testBeamEnd {
		b := Beam{e.beamStart, 0, 33, nil}
		if e.exp != b.Ends(e.walls) {
			t.Errorf("Exp %v got %v", e.exp, b.Ends(e.walls))
		}
	}
}
