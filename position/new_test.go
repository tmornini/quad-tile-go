package position

import "testing"

func TestNew(t *testing.T) {
	New(RandomData())
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(RandomData())
	}
}
