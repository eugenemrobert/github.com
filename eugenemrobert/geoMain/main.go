/*package main

import (
	"fmt"
	"log"

	"github.com/eugenemrobert/geo"
)

func main() {
	coordinates := geo.Coordinates{}
	err := coordinates.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = coordinates.SetLongitude(-1122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(coordinates.Latitude())
	fmt.Println(coordinates.Longitude())
}*/
package main

import (
	"fmt"
	"log"

	"github.com/eugenemrobert/geo"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
