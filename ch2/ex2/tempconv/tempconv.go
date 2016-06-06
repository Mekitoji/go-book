package tempconv

import "fmt"

// Celsius is an unit of measurement for temperature
type Celsius float64

// Fahrenheit is an unit of measurement for temperature
type Fahrenheit float64

// Kelvin is an unit of measurement for temperature
type Kelvin float64

const (
	// AbsoluteZeroC is an absolute zero in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is an temperature of freezing water
	FreezingC Celsius = 0
	// BoilingC is an water boiling temperature
	BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
