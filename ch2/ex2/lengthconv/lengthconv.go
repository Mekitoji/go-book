package lengthconv

import "fmt"

// Feet is an unit of length measurement
type Feet float64

// Meter is an unit of length measurement
type Meter float64

// Inch is an unit of length measurement
type Inch float64

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (i Inch) String() string  { return fmt.Sprintf("%gin", i) }
