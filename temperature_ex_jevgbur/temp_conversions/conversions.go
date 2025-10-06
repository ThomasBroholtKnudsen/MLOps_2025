package temp_conversions

type Temperature struct {
	Value float64
	Unit  string
}

// use pointer *Temperature to modify the original struct
// passing its memory address instead of a copy
func CelsiusToFahrenheit(t *Temperature) *Temperature {
	t.Value = (t.Value * 1.8) + 32
	t.Unit = "F"
	return &Temperature{Value: t.Value, Unit: t.Unit} // creates a new struct , but returns the pointer to it &Temperature..
}

func FahrenheitToCelsius(t *Temperature) *Temperature {
	t.Value = (t.Value - 32) / 1.8
	t.Unit = "C"
	return &Temperature{Value: t.Value, Unit: t.Unit}
}

