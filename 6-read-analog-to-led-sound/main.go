package main

import (
	"embedded-arduino/utils"
	"machine"
)

func main() {
	machine.InitADC()

	ldr := machine.ADC{Pin: machine.ADC0}
	ldr.Configure(machine.ADCConfig{})

	buzzer := machine.Timer1
	buzzer.Configure(machine.PWMConfig{Period: 1})
	buzzerChannel, err := buzzer.Channel(machine.D9)
	if err != nil {
		println("init buzzer channel err : ", err)
	}

	led := machine.Timer2
	led.Configure(machine.PWMConfig{Period: 16384e3})
	ledChannel, err := led.Channel(machine.D11)
	if err != nil {
		println("init led channel : ", err)
	}

	for {
		led.Set(ledChannel, uint32(ldr.Get()/257))
		buzzer.Set(buzzerChannel, uint32(ldr.Get()))
		utils.Delay(200)
	}
}
