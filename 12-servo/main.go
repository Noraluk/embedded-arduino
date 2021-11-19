package main

import (
	"embedded-arduino/utils"
	"machine"
	"time"

	"tinygo.org/x/drivers/servo"
)

var (
	pwm = machine.Timer1
	pin = machine.D9
)

// ref : https://github.com/tinygo-org/drivers/blob/release/examples/servo/servo.go
func main() {
	s, err := servo.New(pwm, pin)
	if err != nil {
		for {
			println("could not configure servo")
			time.Sleep(time.Second)
		}
	}

	for {
		println("setting to 0°")
		s.SetMicroseconds(1000)
		utils.Delay(1000)

		println("setting to 45°")
		s.SetMicroseconds(1500)
		utils.Delay(1000)

		println("setting to 90°")
		s.SetMicroseconds(2000)
		utils.Delay(1000)
	}
}
