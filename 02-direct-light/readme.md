# Light and Shadow in 2D

In the future:  
On a 2D space, drop some light and check its shadow.

## Setup

![canvas](docs/canvas.png)

A rectangle 125x32.  
Beam Sources shooting red light with various angles.

Helper info about canvas, positioning and color pallette in this [doc](https://docs.google.com/spreadsheets/d/1BydYF5Aa_xKUDXgPKSL9cN4QbD8yq2vslvwRQBO_mws/edit?usp=sharing).

## Reasoning

TBD

## run

```bash
# dependencies
go get -u github.com/llgcode/draw2d
```

```bash
# run
go run ./...
open canvas.png
# test
go test ./...
```

## next steps

- Light the wall based on light hitting in
- Render obstacles that will prevent the wall from lighting

Much later...

- Monte Carlo to make it faster!
