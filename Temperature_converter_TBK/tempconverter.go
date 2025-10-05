package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Write your code here

	value, _ := strconv.ParseFloat(os.Args[1], 64)
	unit := os.Args[2]

	var temp *Temperature //Kunne også skrive uden *. Se følgende kommentarer.

	temp = &Temperature{Value: value, Unit: unit} //Hvis jeg skrev uden * i linje 15, så behøver jeg ikke & her.

	var converted *Temperature

	if temp.Unit == "C" {
		converted = celsiusToFahrenheit(temp) // Hvis jeg ikke brugte * og & ovenover, så skulle jeg bruge & foran temp her.
	} else {
		converted = fahrenheitToCelsius(temp)
	}

	fmt.Printf("%.0f %s\n", converted.Value, converted.Unit)
}

func celsiusToFahrenheit(t *Temperature) *Temperature {
	// Implement conversion logic
	converted_temp := &Temperature{Value: (t.Value * 1.8) + 32, Unit: "F"}

	return converted_temp
}

func fahrenheitToCelsius(t *Temperature) *Temperature {
	// Implement conversion logic
	converted_temp := &Temperature{Value: (t.Value - 32) / 1.8, Unit: "C"}

	return converted_temp
}

type Temperature struct {
	Value float64
	Unit  string
}
