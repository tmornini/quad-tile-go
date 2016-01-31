package position

import "math/rand"

func Random() *Position {
	return New(
		RandomData(),
	)
}

func RandomData() (float64, float64, float64) {
	return rand.Float64()*180.0 - 90.0,
		rand.Float64()*360.0 - 180.0,
		rand.Float64()*100000.0 - 1000.0
}
