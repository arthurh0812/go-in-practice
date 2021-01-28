package tempconv

import (
	"flag"
	"fmt"
	"strconv"
)

// Celsius stores a celcius-based temperature value
type Celsius float64

func (c Celsius) String() string {
	v := strconv.FormatFloat(float64(c), 'f', -1, 10)
	return v + "°C"
}

// Fahrenheit stores a fahrenheit-based temperature value
type Fahrenheit float64

func (f Fahrenheit) String() string {
	v := strconv.FormatFloat(float64(f), 'f', -1, 10)
	return v + "°F"
}

// CToF converts the given celsius-based temperature value to the corresponding fahrenheit value
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*1.8 + 32)
}

// FToC converts the given fahrenheit-based temperature value to the corresponding celsius value
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) / 1.8)
}

type celciusFlag struct {
	Celsius
}

func (f *celciusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	default:
		return fmt.Errorf("failed to convert temperature %q", s)
	}
}

// CelciusFlag creates and returns a new command-line flag with the given name, usage and default value
func CelciusFlag(name string, value Celsius, usage string) *Celsius {
	f := &celciusFlag{Celsius: value}
	flag.CommandLine.Var(f, name, usage)
	return &f.Celsius
}

// CelsiusFlagVar creates a new command-line flag with the given name, usage and default value, and stores the result in p
func CelsiusFlagVar(p *Celsius, name string, value Celsius, usage string) {
	p = CelciusFlag(name, value, usage)
}
