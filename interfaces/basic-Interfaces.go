package goInterface

import (
	log "github.com/sirupsen/logrus"
)

type Shape interface {
	Area() int
}

type Circle struct {
	Radius int
}

type Square struct {
	Side int
}

func (s *Square) Area() int {
	return s.Side * s.Side
}

func (c *Circle) Area() int {
	return 2 * c.Radius
}

func BasicInterfaceExample() {
	log.Debug("\n--------------   Go simple example on Interfaces --------------\n")
	sq := Square{4}
	ci := Circle{3}

	var sh Shape

	sh = &sq
	log.Debugf("\nThe calculated Area for Shape ( Square ) is %v", sh.Area())

	sh = &ci
	log.Debugf("\nThe calculated Area for Shape ( Circle ) is %v", sh.Area())

}

func GetArea(s Shape) int {
	return s.Area()
}
