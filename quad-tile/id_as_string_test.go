package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestIdAsString(t *testing.T) {
	pos, _ := position.New(0, -0.000000001, 0)

	if qt := New(pos, 16); qt.IdAsString(2) != "ad" {
		t.Error(fmt.Sprintf("expected: ad, got: %#v\n", qt.IdAsString(2)))
	}

	if qt := New(pos, 16); qt.IdAsString(16) != "addddddddddddddd" {
		t.Error(fmt.Sprintf("expected: addddddddddddddd, got: %#v\n", qt.IdAsString(16)))
	}

	pos, _ = position.New(0, 0, 0)

	if qt := New(pos, 16); qt.IdAsString(2) != "bc" {
		t.Error(fmt.Sprintf("expected: bc, got: %#v\n", qt.IdAsString(2)))
	}

	if qt := New(pos, 16); qt.IdAsString(16) != "bccccccccccccccc" {
		t.Error(fmt.Sprintf("expected: bccccccccccccccc, got: %#v\n", qt.IdAsString(16)))
	}

	pos, _ = position.New(-0.000000001, -0.000000001, 0)

	if qt := New(pos, 16); qt.IdAsString(2) != "cb" {
		t.Error(fmt.Sprintf("expected: cb, got: %#v\n", qt.IdAsString(2)))
	}

	if qt := New(pos, 16); qt.IdAsString(16) != "cbbbbbbbbbbbbbbb" {
		t.Error(fmt.Sprintf("expected: cbbbbbbbbbbbbbbb, got: %#v\n", qt.IdAsString(16)))
	}

	pos, _ = position.New(-0.000000001, 0, 0)

	if qt := New(pos, 16); qt.IdAsString(2) != "da" {
		t.Error(fmt.Sprintf("expected: da, got: %#v\n", qt.IdAsString(2)))
	}

	if qt := New(pos, 16); qt.IdAsString(16) != "daaaaaaaaaaaaaaa" {
		t.Error(fmt.Sprintf("expected: daaaaaaaaaaaaaaa, got: %#v\n", qt.IdAsString(16)))
	}
}

func TestIdAsStringPanicOnLowLevel(t *testing.T) {
	defer func() { recover() }()

	pos, _ := position.New(0, 0, 0)

	New(pos, 16).IdAsUint32(0)

	t.Error("expected panic on level < 1")
}

func TestIdAsStringPanicOnHigherThanQuadTileLevel(t *testing.T) {
	defer func() { recover() }()

	pos, _ := position.New(0, 0, 0)

	New(pos, 8).IdAsUint32(9)

	t.Error("expected panic on level > QuadTile.level")
}

func BenchmarkIdAsString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(position.Random(), 16).IdAsString(16)
	}
}
