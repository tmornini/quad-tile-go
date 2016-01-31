package quadTile

import "testing"

func TestNewLevel(t *testing.T) {
	NewLevel()
}

func BenchmarkNewLevel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewLevel()
	}
}
