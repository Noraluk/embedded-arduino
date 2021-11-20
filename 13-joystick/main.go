package main

import (
	"machine"
)

var (
	pinX = machine.ADC{Pin: machine.ADC0}
	pinY = machine.ADC{Pin: machine.ADC1}
	led  = machine.D13
	btn  = machine.D2
)

const (
	modifyNum = 64
)

func main() {
	setup()
	loop()
}

func setup() {
	machine.InitADC()

	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	btn.Configure(machine.PinConfig{Mode: machine.PinInput})

	pinX.Configure(machine.ADCConfig{})
	pinY.Configure(machine.ADCConfig{})
}

func loop() {
	for {
		x := pinX.Get() / modifyNum
		y := pinY.Get() / modifyNum

		println(x, y)

		btn := btn.Get()
		switch btn {
		case true:
			led.High()
			break
		default:
			led.Low()
			break
		}
	}
}
