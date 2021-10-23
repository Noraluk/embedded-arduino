package main

import (
	"embedded-arduino/utils"
	"machine"
)

func main() {

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		utils.Delay(500)
	}

}
