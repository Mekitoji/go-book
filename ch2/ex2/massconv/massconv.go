package massconv

import "fmt"

// Kilogram is a unit of measurement for mass
type Kilogram float64

// Pound is a unit of measurement for mass
type Pound float64

// Stone is a unit of measurement for mass
type Stone float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }

func (p Pound) String() string { return fmt.Sprintf("%glb", p) }

func (s Stone) String() string { return fmt.Sprintf("%gst", s) }
