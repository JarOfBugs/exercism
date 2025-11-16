package space

type Planet string

const EARTH_YEAR_SECONDS float64 = 31_557_600.0

func Age(seconds float64, planet Planet) float64 {
	var multiplier float64
    switch planet {
        case "Mercury": 
        	multiplier = 0.2408467
        case "Venus":
        	multiplier = 0.61519726
        case "Mars":
        	multiplier = 1.8808158
        case "Jupiter":
        	multiplier = 11.862615
        case "Saturn":
        	multiplier = 29.447498
        case "Uranus":
        	multiplier = 84.016846
        case "Neptune":
        	multiplier = 164.79132
        case "Earth":
        	multiplier = 1.0
        default:
        	multiplier = -1.0
    }
    if multiplier < 0.0 {
        return multiplier
    }
    return seconds / (EARTH_YEAR_SECONDS * multiplier)
}
