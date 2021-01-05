package design_patterns

import (
	"testing"
)

func TestBuilder(t*testing.T)  {
	boss:=new(Boos)

	carbuilder:=new(CarBuilder)
	boss.Hire(carbuilder)

	car:=boss.GenCar("BWM")
	t.Log(car.Brand)
}
