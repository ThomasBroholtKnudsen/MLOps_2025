package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	converter "temperature_converter/conversions"
)

func main() {
	// Expect exactly 2 arguments: <value> <unit>
	if len(os.Args) != 3 {
		fmt.Println("Usage: <binary> <value> <unit>")
		os.Exit(1)
	}

	// Parse numeric value
	value, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid temperature value. Please enter a numeric value.")
		os.Exit(1)
	}

	// Normalize and validate unit
	unit := strings.ToUpper(os.Args[2])
	if unit != "C" && unit != "F" {
		fmt.Println("Invalid unit. Please provide 'C' for Celsius or 'F' for Fahrenheit.")
		os.Exit(1)
	}

	// Build a pointer to Temperature
	t := &converter.Temperature{Value: value, Unit: unit}

	// Call the appropriate conversion
	var result *converter.Temperature
	if t.Unit == "C" {
		result = converter.CelsiusToFahrenheit(t)
	} else {
		result = converter.FahrenheitToCelsius(t)
	}

	// Print without decimal places
	fmt.Printf("%.0f %s\n", result.Value, result.Unit)
}
