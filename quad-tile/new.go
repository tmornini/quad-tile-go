package quadTile

import "github.com/tmornini/quad-tile-go/position"

type QuadTile struct {
	Id uint
  Level uint
	Position *position.Position
}

func New(pos *position.Position, desired_level uint) *QuadTile {

  var id uint
  var latitude = pos.Latitude
  var longitude = pos.Longitude
  var center_latitude float64
  var center_longitude float64
  var level uint = 0
  var latitude_adjustment float64 = 45
  var longitude_adjustment float64 = 90

  for level < desired_level {
    var quadrant uint

    if longitude < center_longitude { // west
      center_longitude -= longitude_adjustment
    } else { // east
      quadrant |= 1
      center_longitude += longitude_adjustment
    }

    if latitude >= center_latitude { // north
      center_latitude += latitude_adjustment
    } else { // south
      quadrant |= 2
      center_latitude -= latitude_adjustment
    }

    id += quadrant << (level * 2)

    latitude_adjustment /= 2
    longitude_adjustment /= 2

    level++
  }

	return &QuadTile{
		id,
    level,
		pos,
  }
}
