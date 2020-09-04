package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32

// function to convert celcius to fahrenheit
func CtoFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((t*9/5)+32)
}

// function to convert fahrenheit to celcius
func FtoCelcius(t Fahrenheit) Celsius {
	return Celsius((t-32)*5/9)
}

func main() {

	var tempFahr Fahrenheit = 100
	var tempCelsius Celsius = 100

	tempFahr = CtoFahrenheit(tempCelsius)
	tempCelsius = FtoCelcius(tempFahr)
	fmt.Printf("%f C is equal to %f F\n", tempCelsius, tempFahr)
	fmt.Printf("%f F is equal to %f C\n", tempFahr, tempCelsius)

}