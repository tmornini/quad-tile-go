package quadTile

func (self *QuadTile) IdAsString() string {
  id := self.Id
  id_string := ""

  for i := uint(0); i < self.Level; i++ {
    id_string += string(97 + (id & 3))
    id = id >> 2
  }

  return id_string
}
