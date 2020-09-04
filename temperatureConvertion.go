package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// function to convert celcius to fahrenheit
func toFahrenheit(t Celsius) Fahrenheit {

	return Fahrenheit((t*9/5)+32)

}

func main() {

	var tempCelsius Celsius = 100
	tempFahr := toFahrenheit(tempCelsius)
	fmt.Printf("%f C is equal to %f F", tempCelsius, tempFahr)

}