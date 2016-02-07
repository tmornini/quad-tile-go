package quadTile

import "fmt"

func (self *QuadTile) IdAsString(level uint) string {
  if level < 1 {
    panic(fmt.Sprintf("level %v (< 1)", level))
  } else if level > self.Level {
    panic(fmt.Sprintf("level (%v) than QuadTile (%v)", level, self.Id))
  }

  id := self.Id
  id_string := ""

  for i := uint(0); i < level; i++ {
    id_string += string(97 + (id >> 62))
    id = id << 2
  }

  return id_string
}
