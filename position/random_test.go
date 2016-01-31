package position

import "testing"

func TestRandom(t *testing.T) {
	Random()
}

func BenchmarkRandom(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Random()
	}
}
