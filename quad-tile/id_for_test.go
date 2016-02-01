package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestIdFor(t *testing.T) {
	if id := IdFor(position.New(0, -0.000000001, 0), 16); id != "addddddddddddddd" {
		t.Error(fmt.Sprintf("expected: addddddddddddddd, got: %#v\n", id))
	}

	if id := IdFor(position.New(0, 0, 0), 16); id != "bccccccccccccccc" {
		t.Error(fmt.Sprintf("expected: bccccccccccccccc, got: %#v\n", id))
	}

	if id := IdFor(position.New(-0.000000001, -0.000000001, 0), 16); id != "cbbbbbbbbbbbbbbb" {
		t.Error(fmt.Sprintf("expected: cbbbbbbbbbbbbbbb, got: %#v\n", id))
	}

	if id := IdFor(position.New(-0.000000001, 0, 0), 16); id != "daaaaaaaaaaaaaaa" {
		t.Error(fmt.Sprintf("expected: daaaaaaaaaaaaaaa, got: %#v\n", id))
	}
}

func BenchmarkIdFor(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IdFor(position.Random(), 16)
	}
}
