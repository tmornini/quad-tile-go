package quadTile

import "fmt"

import "github.com/tmornini/quad-tile-go/position"

type QuadTile struct {
	Id uint64
  Level uint
	Position *position.Position
}

func New(pos *position.Position, level uint) *QuadTile {

  if level < 1 {
    panic(fmt.Sprintf("cannot handle desired level %v (< 1)", level))
  } else if level > 32 {
    panic(fmt.Sprintf("cannot handle desired level %v (> 32)", level))
  }

  var id uint64
  var latitude = pos.Latitude
  var longitude = pos.Longitude
  var center_latitude float64
  var center_longitude float64
  var current_level uint = 0
  var latitude_adjustment float64 = 45
  var longitude_adjustment float64 = 90

  for current_level < level {
    var quadrant uint64

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

    id += quadrant << ((31 - current_level) * 2)

    latitude_adjustment /= 2
    longitude_adjustment /= 2

    current_level++
  }

	return &QuadTile{
		id,
    level,
		pos,
  }
}
