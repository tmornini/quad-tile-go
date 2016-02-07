package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestNew(t *testing.T) {
  pos, _ := position.New(0, -0.000000001, 0)

  if qt := New(pos, 16); qt.Id != 0xfffffffc {
    t.Error(fmt.Sprintf("expected: 0xfffffffc, got: %x\n", qt.Id))
  }

  pos, _ = position.New(0, 0, 0)

  if qt := New(pos, 16); qt.Id != 0xaaaaaaa9 {
    t.Error(fmt.Sprintf("expected: 0xaaaaaaa9, got: %x\n", qt.Id))
  }

  pos, _ = position.New(-0.000000001, -0.000000001, 0)

  if qt := New(pos, 16); qt.Id != 0x55555556 {
    t.Error(fmt.Sprintf("expected: 0x55555556, got: %x\n", qt.Id))
  }

  pos, _ = position.New(-0.000000001, 0, 0)

  if qt := New(pos, 16); qt.Id != 0x00000003 {
    t.Error(fmt.Sprintf("expected: 0x00000003, got: %x\n", qt.Id))
  }
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(position.Random(), 16)
	}
}
