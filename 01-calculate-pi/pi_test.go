package main

import "testing"

var tests = []struct {
	p   Point
	exp bool
}{
	{Point{0, 0}, true},
	{Point{1, 1}, false},
	{Point{0.7, 0.7}, true},
	{Point{0.8, 0.8}, false},
	{Point{0.8, 0.5}, true},
}

func TestIsInCircle(t *testing.T) {
	for _, e := range tests {
		bRet := isInCircle(e.p)
		if bRet != e.exp {
			t.Errorf("Point %f %f expected in circle:%t", e.p.x, e.p.y, e.exp)
		}
	}

}
