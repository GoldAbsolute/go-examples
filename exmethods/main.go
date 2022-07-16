package exmethods

import "fmt"

type planet struct {
	name, color string
}

func (r *planet) greetings() string {
	return "Name this planet is " + r.name + " and this color is " + r.color
}

func Main() {
	newPlanet := planet{"Earth", "Blue"}
	msg := newPlanet.greetings()
	fmt.Println(msg)
	var planet2 = new(planet)
	planet2.name = "Mars"
	planet2.color = "Red"
	msg2 := planet2.greetings()
	fmt.Println(msg2)

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
