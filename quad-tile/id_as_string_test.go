package quadTile

import "fmt"
import "testing"

import "github.com/tmornini/quad-tile-go/position"

func TestIdAsString(t *testing.T) {
  pos, _ := position.New(0, -0.000000001, 0)

  if qt := New(pos, 16); qt.IdAsString() != "addddddddddddddd" {
    t.Error(fmt.Sprintf("expected: addddddddddddddd, got: %#v\n", qt.IdAsString()))
  }

  pos, _ = position.New(0, 0, 0)

  if qt := New(pos, 16); qt.IdAsString() != "bccccccccccccccc" {
    t.Error(fmt.Sprintf("expected: bccccccccccccccc, got: %#v\n", qt.IdAsString()))
  }

  pos, _ = position.New(-0.000000001, -0.000000001, 0)

  if qt := New(pos, 16); qt.IdAsString() != "cbbbbbbbbbbbbbbb" {
    t.Error(fmt.Sprintf("expected: cbbbbbbbbbbbbbbb, got: %#v\n", qt.IdAsString()))
  }

  pos, _ = position.New(-0.000000001, 0, 0)

  if qt := New(pos, 16); qt.IdAsString() != "daaaaaaaaaaaaaaa" {
    t.Error(fmt.Sprintf("expected: daaaaaaaaaaaaaaa, got: %#v\n", qt.IdAsString()))
  }
}

func BenchmarkIdAsString(b *testing.B) {
  for n := 0; n < b.N; n++ {
    New(position.Random(), 16).IdAsString()
  }
}
