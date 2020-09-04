package main

import "fmt"

// aliasing type
type Celsius float32
type Fahrenheit float32
type Kelvin float32

// function to convert celcius to fahrenheit
func CtoFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((t*9/5)+32)
}

// function to convert celcius to kelvin
func CtoKelvin(t Celsius) Kelvin {
	return Kelvin(t+273.15)
}

// function to convert fahrenheit to celcius
func FtoCelcius(t Fahrenheit) Celsius {
	return Celsius((t-32)*5/9)
}

func main() {

	var Fahrenheit Fahrenheit = 100
	var Celsius Celsius = 100

	tempFahr := CtoFahrenheit(Celsius)
	tempKelvin := CtoKelvin(Celsius)
	tempCelsius := FtoCelcius(Fahrenheit)
	fmt.Printf("%.2f C is equal to %.2f F\n", Celsius, tempFahr)
	fmt.Printf("%.2f C is equal to %.2f K\n", Celsius, tempKelvin)
	fmt.Printf("%.2f F is equal to %.2f C\n", Fahrenheit, tempCelsius)

}