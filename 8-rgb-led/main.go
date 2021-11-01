package main

import (
	"embedded-arduino/utils"
	"machine"
)

type color struct {
	red   uint32
	green uint32
	blue  uint32
}

type pin struct {
	red   machine.PWM
	blue  machine.PWM
	green machine.PWM
}

func newPin(redPin, bluePin, greenPin machine.PWM) *pin {
	return &pin{red: redPin, blue: bluePin, green: greenPin}
}

type pinChann struct {
	red   uint8
	blue  uint8
	green uint8
}

func newPinChann(redChann, blueChann, greenChann uint8) *pinChann {
	return &pinChann{red: redChann, blue: blueChann, green: greenChann}
}

type rgb struct {
	pin   *pin
	chann *pinChann
}

func newRGB(pin *pin, chann *pinChann) *rgb {
	return &rgb{pin: pin, chann: chann}
}

func main() {
	redPin := machine.Timer2
	bluePin := machine.Timer1
	greenPin := machine.Timer1

	redPin.Configure(machine.PWMConfig{Period: 16384e3})
	bluePin.Configure(machine.PWMConfig{Period: 10000e1})
	greenPin.Configure(machine.PWMConfig{Period: 100000})

	rpChann, err := redPin.Channel(machine.D11)
	if err != nil {
		println(err)
	}

	bpChann, err := bluePin.Channel(machine.D10)
	if err != nil {
		println(err)
	}

	gpChann, err := greenPin.Channel(machine.D9)
	if err != nil {
		println(err)
	}

	pin := newPin(redPin, bluePin, greenPin)
	pinChann := newPinChann(rpChann, bpChann, gpChann)
	rgb := newRGB(pin, pinChann)

	colors := []color{
		{red: 255, green: 0, blue: 0},     // red
		{red: 0, green: 255, blue: 0},     // green
		{red: 0, green: 0, blue: 255},     // blue
		{red: 255, green: 255, blue: 0},   // yellow
		{red: 255, green: 255, blue: 255}, // white
		{red: 128, green: 0, blue: 255},   // purple
		{red: 199, green: 21, blue: 133},  // mediumvioletred
		{red: 165, green: 42, blue: 42},   // brown
		{red: 0, green: 0, blue: 0},       // off
	}

	for {
		for _, color := range colors {
			rgb.color(color.red, color.green, color.blue)
		}
	}
}

func (e *rgb) color(red, green, blue uint32) {
	e.pin.red.Set(e.chann.red, red)
	e.pin.blue.Set(e.chann.blue, blue)
	e.pin.green.Set(e.chann.green, green)
	utils.Delay(1000)
}
