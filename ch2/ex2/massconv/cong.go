package massconv

// PToK convert Pound to Kilogram
func PToK(p Pound) Kilogram { return Kilogram(p / 2.2046) }

// PToS convert Pound to Stone
func PToS(p Pound) Stone { return Stone(p * 0.071429) }

// KToP convert Kilogram to Pound
func KToP(p Kilogram) Pound { return Pound(p * 2.2046) }

// KToS convert Kilogram to Stone
func KToS(k Kilogram) Stone { return Stone(k * 0.15747) }

// SToP convert Stone to Pound
func SToP(p Stone) Pound { return Pound(p * 14.0) }

// SToK convert Stone to Kilogram
func SToK(st Stone) Kilogram { return Kilogram(st / 0.15747) }
