package tempconv

import "fmt"

// Celsius represent unit of measurement for temperature
type Celsius float64

// Fahrenheit represent unit of measurement for temperature
type Fahrenheit float64

// Kelvin represnet unit of measurement for temperature
type Kelvin float64

const (
	// AbsoluteZeroC represent absolute zero in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC represent temperature of freezing water
	FreezingC Celsius = 0
	// BoilingC represent water boiling temperature
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
