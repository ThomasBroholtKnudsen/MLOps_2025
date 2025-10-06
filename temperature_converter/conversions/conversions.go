package conversions

// Temperature holds a numeric value and unit ("C" or "F").
type Temperature struct {
	Value float64
	Unit  string
}

// CelsiusToFahrenheit converts Celsius -> Fahrenheit
func CelsiusToFahrenheit(t *Temperature) *Temperature {
	return &Temperature{
		Value: (t.Value*9.0/5.0 + 32.0),
		Unit:  "F",
	}
}

// FahrenheitToCelsius converts Fahrenheit -> Celsius
func FahrenheitToCelsius(t *Temperature) *Temperature {
	return &Temperature{
		Value: (t.Value - 32.0) * 5.0 / 9.0,
		Unit:  "C",
	}
}
