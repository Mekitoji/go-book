package main

import (
	"fmt"

	"github.com/mekitoji/go-book/ch2/ex2/lengthconv"
	"github.com/mekitoji/go-book/ch2/ex2/massconv"
	"github.com/mekitoji/go-book/ch2/ex2/tempconv"
)

func tconv(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)

	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), f, tempconv.FToK(f))
	fmt.Printf("%s = %s, %s = %s\n", k, tempconv.KToC(k), k, tempconv.KToF(k))
	fmt.Printf("%s = %s, %s = %s\n\n", c, tempconv.CToK(c), c, tempconv.CToF(c))
}

func lconv(l float64) {
	f := lengthconv.Feet(l)
	m := lengthconv.Meter(l)
	i := lengthconv.Inch(l)

	fmt.Printf("%s = %s, %s = %s\n", f, lengthconv.FToM(f), f, lengthconv.FToI(f))
	fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToF(m), m, lengthconv.MToI(m))
	fmt.Printf("%s = %s, %s = %s\n\n", i, lengthconv.IToM(i), i, lengthconv.IToF(i))
}

func mconv(m float64) {
	k := massconv.Kilogram(m)
	p := massconv.Pound(m)
	s := massconv.Stone(m)

	fmt.Printf("%s = %s, %s = %s\n", k, massconv.KToP(k), k, massconv.KToS(k))
	fmt.Printf("%s = %s, %s = %s\n", p, massconv.PToK(p), p, massconv.PToS(p))
	fmt.Printf("%s = %s, %s = %s\n\n", s, massconv.SToK(s), s, massconv.SToP(s))
}
