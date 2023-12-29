package exercise2_2

import "fmt"

type Meter float64
type Feet float64
type Pounds float64
type Kilograms float64

func (m Meter) String() string     { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string      { return fmt.Sprintf("%g feet", f) }
func (p Pounds) String() string    { return fmt.Sprintf("%g pounds", p) }
func (k Kilograms) String() string { return fmt.Sprintf("%g kilograms", k) }

func MToF(m Meter) Feet {
	return Feet(m * 3.2808)
}

func FToM(f Feet) Meter {
	return Meter(f / 3.2808)
}

func PToK(p Pounds) Kilograms {
	return Kilograms(p / 2.20462)
}

func KToP(k Kilograms) Pounds {
	return Pounds(k * 2.20462)
}
