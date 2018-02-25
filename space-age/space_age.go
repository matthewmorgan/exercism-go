// Package space provides a function for converting
// user's age in Earth-years to other-plant-years
package space

type Planet string

var PLANET_YEARS_PER_EARTH_YEAR = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

const SECONDS_PER_EARTH_YEAR = 31557600

// Age takes a user's age in seconds_age int and return their age in years float on each of sol's
// 8 planets in that's planet's years
func Age(seconds_age float64, planet Planet) float64 {
	return seconds_age / PLANET_YEARS_PER_EARTH_YEAR[planet] / SECONDS_PER_EARTH_YEAR
}
