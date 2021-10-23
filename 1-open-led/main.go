package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledSwitch := true

	for {
		led.Set(ledSwitch)
		ledSwitch = !ledSwitch
		delay(500)
	}

}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
