package quadTile

import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestNew(t *testing.T) {
	New(position.Random())
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(position.Random())
	}
}
