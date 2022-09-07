package main

import (
	"fmt"
	"log"

	"github.com/eugenemrobert/keyboard"
)

func main() {
	fmt.Print("Enter a temperature in FahrenHeit: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f degrees Celsius\n", celsius)
}
