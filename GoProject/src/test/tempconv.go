package main

import (
	"fmt"
)

type Celsius float64

type Fahernhrit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahernhrit {
	return Fahernhrit(c*9/5 + 32)
}

func CFoT(f Fahernhrit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	BoilingC := CToF(BoilingC)
	fmt.Printf("%g\n", BoilingC-CToF(FreezingC))
}
