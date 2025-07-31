// The weather package helps pull information relevant to the users current city
package weather

var CurrentCondition string
var CurrentLocation string

// Returns the current weather condition given a specific city
func Forecast(city, condition string) string {
  // CurrentLocation stores the user's current location,
  // CurrentCondition stores the current weather condition
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
