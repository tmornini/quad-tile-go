package main

import "github.com/tmornini/quad-tile-go/position"
import "github.com/tmornini/quad-tile-go/quad-tile"

type quadTileChannel chan *quadTile.QuadTile

const workers = 16
const iterations = 2000000

func start_worker(channel quadTileChannel, count int) {
	for i := 0; i < count; i++ {
		channel <- quadTile.New(position.Random())
	}
}

func start_workers() (quadTileChannel, int) {
	channel := make(chan *quadTile.QuadTile)

	for i := 0; i < workers; i++ {
		go start_worker(channel, iterations/workers)
	}

	return channel, iterations
}

func main() {
	for i := 0; i < iterations; i++ {
		quadTile.New(position.Random())
	}
}

// func main() {
// 	channel, iterations := start_workers()

// 	for i := 0; i < iterations; i++ {
// 		<- channel
// 	}
// }
