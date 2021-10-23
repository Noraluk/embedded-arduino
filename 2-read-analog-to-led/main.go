package main

import (
	"machine"
	"time"
)

const delayBetweenPeriods = time.Second * 5

func main() {
	machine.InitADC()

	ldr := machine.ADC{Pin: machine.ADC0}
	ldr.Configure(machine.ADCConfig{})

	led := machine.Timer2
	led.Configure(machine.PWMConfig{Period: 16384e3})

	channel, err := led.Channel(machine.D11)
	if err != nil {
		println("error : ", err)
	}

	for {
		led.Set(channel, uint32(ldr.Get()/257))
		delay(100)
	}

}

func delay(t uint32) {
	time.Sleep(time.Duration(1000000 * t))
}
