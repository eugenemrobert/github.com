/*package main

import (
	"fmt"

	"github.com/eugenemrobert/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Street = "123 Oak St"
	subscriber.City = "Omaha"
	subscriber.State = "NE"
	subscriber.PostalCode = "68111"
	fmt.Println("Street: ", subscriber.Street)
	fmt.Println("City: ", subscriber.City)
	fmt.Println("State: ", subscriber.State)
	fmt.Println("Postal Code: ", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Carr"}
	employee.Street = "456 Elm St"
	employee.City = "Portland"
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("Street: ", employee.Street)
	fmt.Println("City: ", employee.City)
	fmt.Println("State: ", employee.State)
	fmt.Println("Postal Code: ", employee.PostalCode)
}*/

package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Println("Speeding up")
}
func (c Car) Brake() {
	fmt.Println("Stopping")
}
func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Println("Speeding up")
}
func (t Truck) Brake() {
	fmt.Println("Stopping")
}
func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}
func (t Truck) LoadCargo(cargo string) {
	fmt.Println("Loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(string)
}

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()
	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCargo("test cargo")
	}
}

func main() {
	TryVehicle(Truck("Fnord F180"))
}

/*func main() {
	var vehicle Vehicle = Car("Toyota Yarvic")
	vehicle.Accelerate()
	vehicle.Steer("left")

	vehicle = Truck("Fnord F180")
	vehicle.Brake()
	vehicle.Steer("right")
}*/
