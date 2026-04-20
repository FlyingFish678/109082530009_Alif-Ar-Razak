package main

import "fmt"

// alias/tipe data baru di tingkat global
type suhu float64

func CelsiusToReamur(celsius suhu) suhu {
	return (4.0 / 5.0) * celsius
}

func CelsiusToFahrenheit(celsius suhu) suhu {
	return (9.0/5.0)*celsius + 32
}

func CelsiusToKelvin(celsius suhu) suhu {
	return celsius + 273.15
}

func main() {
	var celsius suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&celsius)

	reamur := CelsiusToReamur(celsius)
	fahrenheit := CelsiusToFahrenheit(celsius)
	kelvin := CelsiusToKelvin(celsius)

	fmt.Printf("\n%.2f celcius = %.2f reamur\n", celsius, reamur)
	fmt.Printf("%.2f celcius = %.2f fahrenheit\n", celsius, fahrenheit)
	fmt.Printf("%.2f celcius = %.2f kelvin\n", celsius, kelvin)
}
