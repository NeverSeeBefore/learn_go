package main

import "fmt"

// 开氏温度 转 华氏温度
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(k float64) float64 {
	k = k*9.0/5.0 + 32
	return k
}

func main() {

	fmt.Print("233° K is", kelvinToCelsius(233.0), "° C\n")
	fmt.Print("0° K is", celsiusToFahrenheit(kelvinToCelsius(0.0)), "° F\n")
}
