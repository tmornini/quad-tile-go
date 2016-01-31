package quadTile

import "github.com/tmornini/quad-tile-go/position"

func IdFor(position *position.Position) string {
	level := NewLevel()

	level.Position = position

	return recurse(level)
}

func recurse(level *Level) string {
	this_level(level)

	if level.is_finished() {
		return level.Id
	}

	next_level(level)

	return recurse(level)
}

func this_level(level *Level) {
	level.Last = level.A
	level.Middle = between(level.Last, level.D)

	update_if_a(level)
	update_if_b(level)
	update_if_c(level)
	update_if_d(level)
}

func next_level(level *Level) {
	level.Current += 1

	level.A = between(level.Last, level.A)
	level.B = between(level.Last, level.B)
	level.C = between(level.Last, level.C)
	level.D = between(level.Last, level.D)
}

func between(a *position.Position, b *position.Position) *position.Position {
	return position.New(
		(a.Latitude+b.Latitude)/2,
		(a.Longitude+b.Longitude)/2,
		(a.Altitude+b.Altitude)/2,
	)
}

func update_if_a(level *Level) {
	if level.Position.Latitude >= level.Middle.Latitude && level.Position.Longitude < level.Middle.Longitude {
		level.Id += "a"
		level.Last = level.A
	}
}

func update_if_b(level *Level) {
	if level.Position.Latitude >= level.Middle.Latitude && level.Position.Longitude >= level.Middle.Longitude {
		level.Id += "b"
		level.Last = level.B
	}
}

func update_if_c(level *Level) {
	if level.Position.Latitude < level.Middle.Latitude && level.Position.Longitude < level.Middle.Longitude {
		level.Id += "c"
		level.Last = level.C
	}
}

func update_if_d(level *Level) {
	if level.Position.Latitude < level.Middle.Latitude && level.Position.Longitude >= level.Middle.Longitude {
		level.Id += "d"
		level.Last = level.D
	}
}
