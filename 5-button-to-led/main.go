package main

import (
	"machine"
)

func main() {

	led := machine.D11
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D7
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		led.Set(!button.Get())
	}
}
