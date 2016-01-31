package quadTile

import "github.com/tmornini/quad-tile-go/position"

type Level struct {
	Id       string
	Current  int
	Desired  int
	Position *position.Position
	A        *position.Position
	B        *position.Position
	C        *position.Position
	D        *position.Position
	Middle   *position.Position
	Last     *position.Position
}

func (level *Level) is_finished() bool {
	return level.Current == level.Desired
}

func NewLevel() *Level {
	return &Level{
		Id:      "",
		Current: 1,
		Desired: 16,
		A:       position.New(90.0, -180.0, 0),
		B:       position.New(90.0, 180.0, 0),
		C:       position.New(-90.0, -180.0, 0),
		D:       position.New(-90.0, 180.0, 0),
	}
}
