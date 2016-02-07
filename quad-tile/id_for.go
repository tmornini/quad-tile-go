package quadTile

import "github.com/tmornini/quad-tile-go/position"

func IdFor(
	poi *position.Position,
	desired_level uint,
) string {
	return recurse(
		"",
		poi.Latitude,
		poi.Longitude,
		0,
		0,
		desired_level,
		1,
		45,
		90,
	)
}

func recurse(
	id string,
	latitude float64,
	longitude float64,
	center_latitude float64,
	center_longitude float64,
	desired_level uint,
	current_level uint,
	latitude_adjustment float64,
	longitude_adjustment float64,
) string {
	for current_level <= desired_level {
		quadrant := 97

		if longitude < center_longitude { // west
			center_longitude -= longitude_adjustment
		} else { // east
			quadrant++
			center_longitude += longitude_adjustment
		}

		if latitude >= center_latitude { // north
			center_latitude += latitude_adjustment
		} else { // south
			quadrant += 2
			center_latitude -= latitude_adjustment
		}

		id += string(quadrant)

		latitude_adjustment /= 2
		longitude_adjustment /= 2

		current_level++
	}

	return id
}
