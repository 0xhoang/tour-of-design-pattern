package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	builder := &ApplicationBuilder{}
	app := builder.Build("victorian")
	chair := app.CreateChair()

	assert.Equal(t, chair.HasLegs(), true)
	assert.Equal(t, chair.SitOn(), "VictorianChar")

	app1 := builder.Build("modern")
	chair1 := app1.CreateChair()
	assert.Equal(t, chair1.HasLegs(), false)
	assert.Equal(t, chair1.SitOn(), "ModernChar")
}
