package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64

type Meter float64
type Foot float64

type Kilogram float64
type Pound float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (f Foot) String() string {
	return fmt.Sprintf("%g f", f)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g lb", p)
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g kg", k)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func MToF(m Meter) Foot {
	return Foot(m * 3.2808)
}

func FToM(f Foot) Meter {
	return Meter(f * 0.3048)
}

func KToP(k Kilogram) Pound {
	return Pound(k * 2.205)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p * 0.4535)
}
