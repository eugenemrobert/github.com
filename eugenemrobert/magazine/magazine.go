package magazine

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name   string
	Salary float64
	Address
}
