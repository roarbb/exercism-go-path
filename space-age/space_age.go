package space

// Planet represents space object name
type Planet string

// Age will compute age of person on particular planet
// seconds to years conversion
func Age(seconds float64, planet Planet) float64 {
	orbitalPeriods := map[Planet]float64{
		"Earth":   1,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return seconds / (float64(31557600) * orbitalPeriods[planet])
}
