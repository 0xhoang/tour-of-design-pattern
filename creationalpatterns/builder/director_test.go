package builder

import "testing"
import "github.com/stretchr/testify/assert"

func TestBuilder(t *testing.T) {
	d := NewDirector()

	carBuilder := NewCarBuilder()
	d.makeSportCar(carBuilder)
	car := carBuilder.GetCar()

	assert.Equal(t, car.Engine, "Sport")
	assert.Equal(t, car.Seats, 2)

	carManualBuilder := NewCarManualBuilder()
	d.makeSuv(carManualBuilder)
	carManual := carManualBuilder.GetCar()

	assert.Equal(t, carManual.Engine, "SUV")
	assert.Equal(t, carManual.Seats, 4)
}
