package quadTile

import "fmt"

func (self *QuadTile) IdAsUint32(level uint) uint32 {

	if level < 1 {
		panic(fmt.Sprintf("level %v (< 1)", level))
	} else if level > 16 {
		panic(fmt.Sprintf("level %v (> 16)", level))
	} else if level > self.Level {
		panic(fmt.Sprintf("level (%v) greater than QuadTile (%v)", level, self.Level))
	}

	return uint32(self.Id >> 32)
}
