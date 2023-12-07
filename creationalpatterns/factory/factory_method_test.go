package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	var app Application
	app.initialize("road")

	assert.Equal(t, app.L.PlanDelivery(), "road")
	transport1 := app.L.CreateTransport()
	assert.Equal(t, transport1.Delivery(), "truck")

	app.initialize("sea")
	assert.Equal(t, app.L.PlanDelivery(), "sea")
	transport2 := app.L.CreateTransport()
	assert.Equal(t, transport2.Delivery(), "ship")
}
