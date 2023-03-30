package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (val celsius) fahrenheit() fahrenheit {
	val = val*9.0/5.0 + 32
	return fahrenheit(val)
}

func (val celsius) kelvin() kelvin {
	val += 263.15
	return kelvin(val)
}

func (val fahrenheit) celsius() celsius {
	val = (val - 32) * 5.0 / 9.0
	return celsius(val)
}

func (val fahrenheit) kelvin() kelvin {
	c := fahrenheit.celsius(val)
	k := celsius.kelvin(c)
	return k
}

func (val kelvin) celsius() celsius {
	val -= 273.15
	c := celsius(val)
	return c
}

func (val kelvin) fahrenheit() fahrenheit {
	c := kelvin.celsius(val)
	f := celsius.fahrenheit(c)
	return f
}

func main() {

	// 233° K
	fmt.Print("233 °K to celsius ", kelvin.celsius(233.0), " °C\n")
	fmt.Print("233 °K is fahrenheit ", kelvin.fahrenheit(233.0), " °F\n")
	fmt.Print("-40.15 °C to kelvin ", celsius.kelvin(-40.15), " °K\n")
	fmt.Print("-40.15 °C to fahrenheit ", celsius.fahrenheit(-40.15), " °F\n")
	fmt.Print("-40.27 °F to celsius ", fahrenheit.celsius(-40.27), " °C\n")
	fmt.Print("-40.27 °F to fahrenheit ", fahrenheit.kelvin(-40.27), " °K\n")
	fmt.Println()
	fmt.Print("0 °K to celsius ", kelvin.celsius(0.0), " °C\n")
	fmt.Print("0 °K is fahrenheit ", kelvin.fahrenheit(0.0), " °F\n")
	fmt.Print("0 °C to kelvin ", celsius.kelvin(0.0), " °K\n")
	fmt.Print("0 °C to fahrenheit ", celsius.fahrenheit(0.0), " °F\n")
	fmt.Print("0 °F to celsius ", fahrenheit.celsius(0.0), " °C\n")
	fmt.Print("0 °F to fahrenheit ", fahrenheit.kelvin(0.0), " °K\n")
}
