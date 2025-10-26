// Package weather provide information about climatic status upon different locations in the world.
package weather

var (
    // CurrentCondition represents the a climatic condition.
	CurrentCondition string
    // CurrentLocation represent a city in the world.
	CurrentLocation  string
)
// Forecast return a string that explains a current weather condition in a location that are given in input.
// The returned string format is <CurrentLocation> - current weather condition: <CurrentCondition>.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
