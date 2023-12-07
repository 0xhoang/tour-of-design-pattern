package main

import "fmt"

type AbstractChar struct {
}

func (a *AbstractChar) CallCommon() string {
	return "call common"
}

//---

type Chair interface {
	HasLegs() bool
	SitOn() string
}

//---

type VictorianChar struct {
	AbstractChar
}

func (v *VictorianChar) HasLegs() bool {
	return true
}

func (v *VictorianChar) SitOn() string {
	fmt.Println(v.CallCommon())
	return "VictorianChar"
}

//---

type ModernChar struct {
	AbstractChar
}

func (m *ModernChar) HasLegs() bool {
	return false
}

func (m *ModernChar) SitOn() string {
	fmt.Println(m.CallCommon())
	return "ModernChar"
}

//---

type FurnitureFactory interface {
	CreateChair() Chair
}

//---

type VictorianFurnitureFactory struct {
}

func (v *VictorianFurnitureFactory) CreateChair() Chair {
	abChar := &AbstractChar{}

	victorChar := &VictorianChar{
		AbstractChar: *abChar,
	}

	return victorChar
}

//---

type ModernFurnitureFactory struct {
}

func (m *ModernFurnitureFactory) CreateChair() Chair {
	abChar := &AbstractChar{}

	modernChar := &ModernChar{
		AbstractChar: *abChar,
	}

	return modernChar
}

//---

type Application struct {
	F FurnitureFactory
}

func (a *Application) Init(f FurnitureFactory) {
	a.F = f

}

func (a *Application) CreateChair() Chair {
	return a.F.CreateChair()
}

type ApplicationBuilder struct {
}

func (a *ApplicationBuilder) Build(types string) Application {
	var factory FurnitureFactory

	switch types {
	case "victorian":
		factory = &VictorianFurnitureFactory{}

	case "modern":
		factory = &ModernFurnitureFactory{}
	}

	app := Application{}
	app.Init(factory)

	return app
}
