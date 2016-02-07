[![Build Status](https://travis-ci.org/tmornini/quad-tile-go.svg?branch=master)](https://travis-ci.org/tmornini/quad-tile-go)

These two Go packages implement the QuadTile hierarchichal binning algorithm described [by the OpenStreetMap project](http://wiki.openstreetmap.org/wiki/QuadTiles).

* github.com/tmornini/quad-tile-go/position
  * New(latitude, longitude, altitude float64) (position, error)
    • returns an error if co-ordinates are out of bounds

* github.com/tmornini/quad-tile-go/quad-tile
  * method New(*position Position.position, level uint)
    • returns a *quadTile
  * method IdAsString(level) string
    • returns Id as a string
  * method IdAsUint32(level) uint32
    • returns Id as a uint32
