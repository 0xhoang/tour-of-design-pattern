package builder

type Director struct {
}

func NewDirector() *Director {
	return &Director{}
}

func (d *Director) makeSuv(builder InterfaceBuilder) {
	builder.reset()
	builder.setSeats(4)
	builder.setEngine("SUV")
	builder.setTripComputer()
	builder.setGPS()
}

func (d *Director) makeSportCar(builder InterfaceBuilder) {
	builder.reset()
	builder.setSeats(2)
	builder.setEngine("Sport")
	builder.setTripComputer()
	builder.setGPS()
}
