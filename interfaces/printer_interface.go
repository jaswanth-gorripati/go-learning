package goInterface

import (
	log "github.com/sirupsen/logrus"
)

type Printer interface {
	Print()
}

type TextPrinter struct {
	Text string
}

func (txt *TextPrinter) Print() {
	log.Debugf("Printing from the Text Interface : %v", txt.Text)
}

func PrintTextFromInterface(p Printer) {
	p.Print()
}
