package tempconv

import (
	"fmt"
)

type Farenheit float64
type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	ZeroC         Celsius = 0
	BoillingC     Celsius = 100
)

func (c Celsius) String() string   { return fmt.Sprintf("%.2f°C", c) }
func (f Farenheit) String() string { return fmt.Sprintf("%.2f°F", f) }
