package quadTile

import "github.com/tmornini/quad-tile-go/position"

func IdFor(
	poi *position.Position,
	desired_level uint,
) string {
	return recurse(
		"",
		poi,
		position.New(0.0, 0.0, 0.0),
		desired_level,
		1,
		45,
		90,
	)
}

func recurse(
	id string,
	poi *position.Position,
	center *position.Position,
	desired_level uint,
	current_level uint,
	latitude_adjustment float64,
	longitude_adjustment float64,
) string {
	if poi.Latitude >= center.Latitude && poi.Longitude < center.Longitude {
		id += "a"
		center = position.New(center.Latitude+latitude_adjustment, center.Longitude-longitude_adjustment, 0)
	} else if poi.Latitude >= center.Latitude && poi.Longitude >= center.Longitude {
		id += "b"
		center = position.New(center.Latitude+latitude_adjustment, center.Longitude+longitude_adjustment, 0)
	} else if poi.Latitude < center.Latitude && poi.Longitude < center.Longitude {
		id += "c"
		center = position.New(center.Latitude-latitude_adjustment, center.Longitude-longitude_adjustment, 0)
	} else if poi.Latitude < center.Latitude && poi.Longitude >= center.Longitude {
		id += "d"
		center = position.New(center.Latitude-latitude_adjustment, center.Longitude+longitude_adjustment, 0)
	} else {
		panic("this should never happen")
	}

	if current_level == desired_level {
		return id
	} else {
		return recurse(
			id,
			poi,
			center,
			desired_level,
			current_level+1,
			latitude_adjustment/2,
			longitude_adjustment/2,
		)
	}
}
