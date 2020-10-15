package main

import (
	"fmt"
	"math"
	"math/rand"
)

// R the radius of the circle
const R = 1

// N total number of shots for the experiment
var N int64 = 100000000

// K number of shots in the circle
var K int64 = 0

// Point in 2 dimensions
type Point struct {
	x float64
	y float64
}

func main() {
	var p Point
	var bRet bool
	var i int64
	var myPi float64

	for i = 0; i < N; i++ {
		p = Point{
			rand.Float64(),
			rand.Float64(),
		}
		bRet = isInCircle(p)
		if bRet {
			K++
		}
	}

	myPi = float64(4*K) / float64(N)
	diffRate := math.Abs(myPi-math.Pi) / math.Pi

	fmt.Println("****************************************************")
	fmt.Println("************  Monte carlo simulation of Pi *********")
	fmt.Println("****************************************************")
	fmt.Printf("Shots:         %20d\n", N)
	fmt.Printf("In Circle:     %20d\n", K)
	fmt.Printf("Pi approx %25f\n", myPi)
	fmt.Printf("Pi known  %25f\n", math.Pi)
	fmt.Printf("Differentiation %20.6f %%\n", 100*diffRate)
	fmt.Println("****************************************************")
}

func isInCircle(p Point) bool {
	return p.x*p.x+p.y*p.y < R*R
}
