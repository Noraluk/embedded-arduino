package main

import (
	"embedded-arduino/utils"
	"machine"
)

func main() {
	machine.InitADC()

	soundSensor := machine.ADC{Pin: machine.ADC0}
	soundSensor.Configure(machine.ADCConfig{})

	led := machine.D13
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	soundMax := 2944
	soundMin := 2880
	soundAvg := (soundMin + soundMax) / 2
	threshold := 3100
	isOn := false

	for {
		sensorVal := soundSensor.Get()

		if sensorVal >= uint16(soundMin) && sensorVal <= uint16(soundMax) {
			sensorVal = uint16(soundAvg)
		}

		if sensorVal >= uint16(threshold) {
			isOn = !isOn
		}
		led.Set(isOn)
		utils.Delay(50)
	}
}
