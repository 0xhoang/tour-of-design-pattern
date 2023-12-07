package builder

type CarManualBuilder struct {
	manual *CarManual
}

func NewCarManualBuilder() *CarManualBuilder {
	return &CarManualBuilder{}
}

func (c *CarManualBuilder) GetCar() *CarManual {
	return c.manual
}

func (c *CarManualBuilder) reset() {
	c.manual = &CarManual{}
}

func (c *CarManualBuilder) setSeats(number int) {
	c.manual.Seats = number
}

func (c *CarManualBuilder) setEngine(engine string) {
	c.manual.Engine = engine
}

func (c *CarManualBuilder) setTripComputer() {
	c.manual.TripComputer = true
}

func (c *CarManualBuilder) setGPS() {
	c.manual.GPS = true
}
