package main

import (
	"embedded-arduino/utils"
	"machine"
)

func main() {
	buzzer := machine.Timer1
	buzzer.Configure(machine.PWMConfig{Period: 1})
	channel, err := buzzer.Channel(machine.D9)
	if err != nil {
		println("init channel err : ", err)
	}

	for {
		buzzer.Set(channel, 200)
		utils.Delay(200)
		buzzer.Set(channel, 0)
		utils.Delay(200)
	}
}
