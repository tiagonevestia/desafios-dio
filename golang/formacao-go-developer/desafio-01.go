package main

import (
	"fmt"
)

func kelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

func main() {
	var kelvin float64

	fmt.Println("Digite a temperatura em Kelvin:")
	fmt.Scan(&kelvin)

	celsius := kelvinToCelsius(kelvin)
	fmt.Printf("A temperatura em Celsius é: %.2f\n", celsius)
}
