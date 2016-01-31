package quadTile

import "github.com/tmornini/quad-tile-go/position"

type QuadTile struct {
	QuadTileId string
	Latitude   float64
	Longitude  float64
	Altitude   float64
}

func New(position *position.Position) *QuadTile {
	return &QuadTile{
		IdFor(position),
		position.Latitude,
		position.Longitude,
		position.Altitude}
}
