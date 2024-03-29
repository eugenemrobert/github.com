package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}
func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals to %0.3f Gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.3f milliliters equals to %0.3f Gallons\n", water, water.ToGallons())
}
