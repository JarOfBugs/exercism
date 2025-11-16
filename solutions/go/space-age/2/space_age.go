package space

type Planet string

const EARTH_YEAR_SECONDS float64 = 31_557_600.0

var planetsWeight = map[Planet]float64 {
    "Mercury" : 0.2408467,
    "Venus"   : 0.61519726,
    "Mars"    : 1.8808158,
    "Jupiter" : 11.862615,
    "Earth"   : 1.0,
    "Saturn"  : 29.447498,
    "Uranus"  : 84.016846,
    "Neptune" : 164.79132,
}

// Age returns given age in years on the given planet
// returns the weight or -1.0 when the given planet is not valid
func Age(seconds float64, planet Planet) float64 {
    multiplier, ok := planetsWeight[planet]
    if !ok {
        return -1.0
    }
    return seconds / (EARTH_YEAR_SECONDS * multiplier)
}
