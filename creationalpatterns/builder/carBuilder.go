package builder

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{}
}

func (c *CarBuilder) GetCar() *Car {
	return c.car
}

func (c *CarBuilder) reset() {
	c.car = &Car{}
}

func (c *CarBuilder) setSeats(number int) {
	c.car.Seats = number
}

func (c *CarBuilder) setEngine(engine string) {
	c.car.Engine = engine
}

func (c *CarBuilder) setTripComputer() {
	c.car.TripComputer = true
}

func (c *CarBuilder) setGPS() {
	c.car.GPS = true
}
