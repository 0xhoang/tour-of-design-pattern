package builder

type InterfaceBuilder interface {
	reset()
	setSeats(number int)
	setEngine(engine string)
	setTripComputer()
	setGPS()
}
