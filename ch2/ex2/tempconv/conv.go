package tempconv

// CToF convert Celsius degree to Fahrenheit degree
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK convert Celsius degree to Kelvin degree
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToC convert Fahrenheit degree to Celsius degree
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK convert Fahrenheit degree to Kelvin degree
func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }

// KToC convert Kelvin degree to Celsius degree
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToF convert Kelvin degree to Fahrenheit degree
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5) - 459.67 }
