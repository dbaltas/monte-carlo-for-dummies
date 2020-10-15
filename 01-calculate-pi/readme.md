# Calculate Pi

Calculating the number Pi (π) using the monte carld simulation technique

## Setup

### Circle

- radius R
- area pi*R*R and
- center on 0.0

### Rectangle

- length 2*R
- area 4*R*R and
- center on 0.0

## Reasoning

Foreach Run there will be `N` number of Shots.

On each Shot, a random point will be created within the rectangle,  
and using the Pythagorean Theorem it will be evaluated whether this point is within the circle or not.  
At the end of the Run we consider `K` the number of shots that had a point in the circle.

The ratio of circle area to rectangle area is `pi/4`   
The ratio of `K/N` should be as close to that `pi/4` as our experiment gets more accurate.  

## run

```
go run 01-calculate-pi/pi.go
****************************************************
************  Monte carlo simulation of Pi *********
****************************************************
Shots:                    100000000
In Circle:                 78537953
Pi approx                  3.141518
Pi known                   3.141593
Differentiation             0.002372 %
****************************************************
```

## next steps

- benchmark tests for multiple runs
- ensure optimized cpu utilization
