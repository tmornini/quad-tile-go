package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestIdAsUint32(t *testing.T) {
  pos, _ := position.New(0, -0.000000001, 0)

  if qt := New(pos, 1); qt.IdAsUint32(1) != 0x00000000 {
    t.Error(fmt.Sprintf("expected: 0x00000000, got: %x\n", qt.IdAsUint32(1)))
  }

  if qt := New(pos, 16); qt.IdAsUint32(16) != 0x3fffffff {
    t.Error(fmt.Sprintf("expected: 0x3fffffff, got: %x\n", qt.IdAsUint32(16)))
  }

  pos, _ = position.New(0, 0, 0)

  if qt := New(pos, 1); qt.IdAsUint32(1) != 0x40000000 {
    t.Error(fmt.Sprintf("expected: 0x40000000, got: %x\n", qt.IdAsUint32(1)))
  }

  if qt := New(pos, 16); qt.IdAsUint32(16) != 0x6aaaaaaa {
    t.Error(fmt.Sprintf("expected: 0x6aaaaaaa, got: %x\n", qt.IdAsUint32(16)))
  }

  pos, _ = position.New(-0.000000001, -0.000000001, 0)

  if qt := New(pos, 1); qt.IdAsUint32(1) != 0x80000000 {
    t.Error(fmt.Sprintf("expected: 0x80000000, got: %x\n", qt.IdAsUint32(1)))
  }

  if qt := New(pos, 16); qt.IdAsUint32(16) != 0x95555555 {
    t.Error(fmt.Sprintf("expected: 0x95555555, got: %x\n", qt.IdAsUint32(16)))
  }

  pos, _ = position.New(-0.000000001, 0, 0)

  if qt := New(pos, 1); qt.IdAsUint32(1) != 0xc0000000 {
    t.Error(fmt.Sprintf("expected: 0xc0000000, got: %x\n", qt.IdAsUint32(1)))
  }

  if qt := New(pos, 16); qt.IdAsUint32(16) != 0xc0000000 {
    t.Error(fmt.Sprintf("expected: 0xc0000000, got: %x\n", qt.IdAsUint32(16)))
  }
}

func TestIdAsUint32PanicOnLowLevel(t *testing.T) {
  defer func() { recover() }()

  pos, _ := position.New(0, 0, 0)

  New(pos, 16).IdAsUint32(0)

  t.Error("expected panic on level < 1")
}

func TestIdAsUint32PanicOnHighLevel(t *testing.T) {
  defer func() { recover() }()

  pos, _ := position.New(0, 0, 0)

  New(pos, 16).IdAsUint32(17)

  t.Error("expected panic on level > 16")
}

func TestIdAsUint32PanicOnHigherThanQuadTileLevel(t *testing.T) {
  defer func() { recover() }()

  pos, _ := position.New(0, 0, 0)

  New(pos, 8).IdAsUint32(9)

  t.Error("expected panic on level > QuadTile.level")
}

func BenchmarkIdAsUint32(b *testing.B) {
  for n := 0; n < b.N; n++ {
    New(position.Random(), 16).IdAsUint32(16)
  }
}
