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

// function to convert fahrenheit to celsius
func FtoCelcius(t Fahrenheit) Celsius {
	return Celsius((t-32)*5/9)
}

// Function to convert fahrenheit to kelvin
func FtoKelvin(t Fahrenheit) Kelvin {
	return Kelvin((t-32)*5/9+273.15)
}

// function to convert kelvin to fahrenheit
func KtoFahrenheit(t Kelvin) Fahrenheit {
	return Fahrenheit((t-273.15)*9/5+32)
}

func KtoCelsius(t Kelvin) Celsius {
	return Celsius(t-273.15)
}

func main() {

	var Fahrenheit Fahrenheit = 100
	var Celsius Celsius = 100
	var Kelvin Kelvin = 100

	/*
		comment before/after line if you want to uncomment set convertion
	*/

	// set fahrenheit convertion from
	tempFahr := CtoFahrenheit(Celsius)
	fmt.Printf("%.2f C is equal to %.2f F\n", Celsius, tempFahr)
	//tempFahr := KtoFahrenheit(Kelvin)
	//fmt.Printf("%.2f K is equal to %.2f F\n", Kelvin, tempFahr)

	// set celsius convertion from
	//tempCelsius := FtoCelcius(Fahrenheit)
	//fmt.Printf("%.2f F is equal to %.2f C\n", Fahrenheit, tempCelsius)
	tempCelsius := KtoCelsius(Kelvin)
	fmt.Printf("%.2f K is equal to %.2f C\n", Kelvin, tempCelsius)

	// set kelvin convertion from
	//tempKelvin := CtoKelvin(Celsius)
	//fmt.Printf("%.2f C is equal to %.2f K\n", Celsius, tempKelvin)
	tempKelvin := FtoKelvin(Fahrenheit)
	fmt.Printf("%.2f F is equal to %.2f K\n", Fahrenheit, tempKelvin)

}