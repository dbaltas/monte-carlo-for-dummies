package main

import (
	"image"
	"math"
)

// Line represents a line between two points
type Line struct {
	a image.Point
	b image.Point
}

// Angle in radians between a source and a target points. Vector angle
func Angle(source image.Point, target image.Point) float64 {
	return math.Atan2((float64)(target.Y-source.Y), (float64)(target.X-source.X))
}

// Intersection finds the intersection of the two lines or nil,
// if the lines are collinear will return nil
// from https://github.com/paulmach/go.geo/blob/master/line.go
func (l Line) Intersection(line Line) *image.Point {
	den := (line.b.Y-line.a.Y)*(l.b.X-l.a.X) - (line.b.X-line.a.X)*(l.b.Y-l.a.Y)
	U1 := (line.b.X-line.a.X)*(l.a.Y-line.a.Y) - (line.b.Y-line.a.Y)*(l.a.X-line.a.X)
	U2 := (l.b.X-l.a.X)*(l.a.Y-line.a.Y) - (l.b.Y-l.a.Y)*(l.a.X-line.a.X)

	if den == 0 {
		// collinear, all bets are off
		if U1 == 0 && U2 == 0 {
			// return InfinityPoint
			return nil
		}

		return nil
	}

	if float64(U1)/float64(den) < 0 || float64(U1)/float64(den) > 1 || float64(U2)/float64(den) < 0 || float64(U2)/float64(den) > 1 {
		return nil
	}

	// return U1 / den
	return l.Interpolate((float64)(U1) / (float64)(den))
	// return nil
}

// Interpolate performs a simple linear interpolation, from A to B.
// This function is the opposite of Project.
// from https://github.com/paulmach/go.geo/blob/master/line.go
func (l *Line) Interpolate(percent float64) *image.Point {
	x := (int)(percent * (float64)(l.b.X-l.a.X))
	y := (int)(percent * (float64)(l.b.Y-l.a.Y))
	return &image.Point{
		l.a.X + x,
		l.a.Y + y,
	}
}

// Distance between two points
func Distance(p1, p2 image.Point) float64 {
	x := float64((p1.X - p2.X))
	y := float64((p1.Y - p2.Y))
	return math.Sqrt(x*x + y*y)
}

// DegToRad Angle Measurment Conversion
func DegToRad(v float64) float64 {
	return v * math.Pi / 180.0
}

// RadToDeg Angle Measurment Conversion
func RadToDeg(v float64) float64 {
	return v * 180 / math.Pi
}

// PointInDirection Point based on point, angle and distance
func PointInDirection(p image.Point, rad float64, distance float64) image.Point {
	return image.Point{
		X: int(distance*math.Cos(rad)) + p.X,
		Y: int(distance*math.Sin(rad)) + p.Y,
	}
}
