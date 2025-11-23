package darts

import "math"

// Score calculate the point obtained given a dart shot
// If the shot miss the target it returns 0 point
// If the shot hit the outer circle it returns 1 point
// If the shot hit the middle circle it returns 5 points
// If the bullseye is hit, 10 points are returned
func Score(x, y float64) int {
	distance := math.Sqrt(x * x + y * y)
    switch {
        case distance > 10:
        	return 0
        case distance > 5 && distance <= 10:
        	return 1
        case distance > 1 && distance <= 5:
        	return 5
        default:
        	return 10
    }
}
