package main

import "github.com/tmornini/quad-tile-go/position"
import "github.com/tmornini/quad-tile-go/quad-tile"

const iterations = 10000000

func main() {
	for i := 0; i < iterations; i++ {
		quadTile.New(position.Random(), 16).IdAsString(16)
	}
}
