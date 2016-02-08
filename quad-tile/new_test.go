package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestNew(t *testing.T) {
	pos, _ := position.New(0, -0.000000001, 0)

	if qt := New(pos, 1); qt.Id != 0x0000000000000000 {
		t.Error(fmt.Sprintf("expected: 0x0000000000000000, got: %x\n", qt.Id))
	}

	if qt := New(pos, 32); qt.Id != 0x3fffffffffffffff {
		t.Error(fmt.Sprintf("expected: 0x3fffffffffffffff, got: %x\n", qt.Id))
	}

	pos, _ = position.New(0, 0, 0)

	if qt := New(pos, 1); qt.Id != 0x4000000000000000 {
		t.Error(fmt.Sprintf("expected: 0x4000000000000000, got: %x\n", qt.Id))
	}

	if qt := New(pos, 32); qt.Id != 0x6aaaaaaaaaaaaaaa {
		t.Error(fmt.Sprintf("expected: 0x6aaaaaaaaaaaaaaa, got: %x\n", qt.Id))
	}

	pos, _ = position.New(-0.000000001, -0.000000001, 0)

	if qt := New(pos, 1); qt.Id != 0x8000000000000000 {
		t.Error(fmt.Sprintf("expected: 0x8000000000000000, got: %x\n", qt.Id))
	}

	if qt := New(pos, 32); qt.Id != 0x9555555555555555 {
		t.Error(fmt.Sprintf("expected: 0x9555555555555555, got: %x\n", qt.Id))
	}

	pos, _ = position.New(-0.000000001, 0, 0)

	if qt := New(pos, 1); qt.Id != 0xc000000000000000 {
		t.Error(fmt.Sprintf("expected: 0xc000000000000000, got: %x\n", qt.Id))
	}

	if qt := New(pos, 16); qt.Id != 0xc000000000000000 {
		t.Error(fmt.Sprintf("expected: 0x00000003, got: %x\n", qt.Id))
	}
}

func TestNewPanicOnLowLevel(t *testing.T) {
	defer func() { recover() }()

	pos, _ := position.New(0, 0, 0)

	New(pos, 0)

	t.Error("expected panic on level < 1")
}

func TestNewPanicOnHighLevel(t *testing.T) {
	defer func() { recover() }()

	pos, _ := position.New(0, 0, 0)

	New(pos, 33)

	t.Error("expected panic on level > 32")
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(position.Random(), 16)
	}
}
