package main

import "fmt"

type Logistics interface {
	PlanDelivery() string
	CreateTransport() Transport
}

//---

type RoadLogistics struct {
}

func (r *RoadLogistics) PlanDelivery() string {
	return "road"
}

func (r *RoadLogistics) CreateTransport() Transport {
	return &Truck{}
}

//---

type SeaLogistics struct {
}

func (r *SeaLogistics) PlanDelivery() string {
	return "sea"
}

func (r *SeaLogistics) CreateTransport() Transport {
	return &Ship{}
}

//---

type Transport interface {
	Delivery() string
}

type Truck struct {
}

func (t *Truck) Delivery() string {
	return "truck"
}

//---

type Ship struct {
}

func (t *Ship) Delivery() string {
	return "ship"
}

//---

type Application struct {
	L Logistics
}

func (a *Application) initialize(transport string) {
	var l Logistics

	switch transport {
	case "road":
		l = &RoadLogistics{}
	case "sea":
		l = &SeaLogistics{}
	default:
		panic("invalid transport")
	}

	a.L = l
}

func main() {
	var app Application
	app.initialize("road")

	fmt.Println("choose plan delivery: ", app.L.PlanDelivery())
	transport1 := app.L.CreateTransport()
	fmt.Println("choose delivery method: ", transport1.Delivery())

	app.initialize("sea")
	fmt.Println("choose plan delivery: ", app.L.PlanDelivery())
	transport2 := app.L.CreateTransport()
	fmt.Println("choose delivery method: ", transport2.Delivery())
}
