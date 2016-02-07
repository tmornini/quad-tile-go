package position

import "math/rand"

func Random() *Position {
	pos, _ := New(RandomData())

	return pos
}

func RandomData() (float64, float64, float64) {
	return rand.Float64()*180 - 90,
		rand.Float64()*359.999999999 - 180,
		rand.Float64()*100000 - 1000
}
