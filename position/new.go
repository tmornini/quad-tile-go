package position

import "fmt"

type Position struct {
	Latitude  float64
	Longitude float64
	Altitude  float64
}

func New(latitude float64, longitude float64, altitude float64) (*Position, error) {
	if latitude < -90 || latitude > 90 || longitude <= -180 || longitude > 180 {
		return nil,
			fmt.Errorf(
				"latitude: %f, longitude: %f, altitude: %f is out of range!",
				latitude,
				longitude,
				altitude,
			)
	}

	return &Position{latitude, longitude, altitude}, nil
}
