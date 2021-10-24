package main

import (
	"machine"
)

func main() {
	machine.InitADC()

	tilt := machine.ADC{Pin: machine.ADC5}
	tilt.Configure(machine.ADCConfig{})

	led := machine.D8
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		switch tiltValue := tilt.Get(); {
		case tiltValue > 200:
			led.Set(true)
			break
		default:
			led.Set(false)
			break
		}
	}
}
