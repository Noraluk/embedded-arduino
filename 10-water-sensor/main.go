package main

import (
	"embedded-arduino/utils"
	"machine"
)

const (
	sensorPower = machine.D7
)

func main() {
	machine.InitADC()

	sensorPower.Configure(machine.PinConfig{Mode: machine.PinOutput})

	sensorPin := machine.ADC{Pin: machine.ADC0}
	sensorPin.Configure(machine.ADCConfig{})

	for {
		sensorPower.Set(true)
		utils.Delay(100)

		sensorValue := sensorPin.Get()
		println(sensorValue)

		sensorPower.Set(false)

		utils.Delay(1000)
	}
}
