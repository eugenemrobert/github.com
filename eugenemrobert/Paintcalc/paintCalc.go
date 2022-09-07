package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("A width of %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("A height of %.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil

}

func main() {
	var amount, total float64
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f litres needed\n", amount)
	total += amount
	fmt.Printf("Total: %.2f litres\n", total)

}
