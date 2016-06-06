package lengthconv

// FToM convert Feet to Meter
func FToM(f Feet) Meter { return Meter(f / 3.2808) }

// FToI convert Feet to Inch
func FToI(f Feet) Inch { return Inch(f * 12.0) }

// MToF convert Meter to Feet
func MToF(m Meter) Feet { return Feet(m * 3.2808) }

// MToI convert Meter to Inch
func MToI(m Meter) Inch { return Inch(m * 39.370) }

// IToM convert Inch to Meter
func IToM(i Inch) Meter { return Meter(i * 0.0254) }

// IToF convert Inch to Feet
func IToF(i Inch) Feet { return Feet(i * 0.083333) }
