package main

import (
	"machine"

	"tinygo.org/x/drivers/keypad4x4"
)

func main() {
	mapping := map[uint8]string{
		0:  "1",
		1:  "2",
		2:  "3",
		3:  "A",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "B",
		8:  "7",
		9:  "8",
		10: "9",
		11: "C",
		12: "*",
		13: "0",
		14: "#",
		15: "D",
	}

	keypadDevice := keypad4x4.NewDevice(machine.D9, machine.D8, machine.D7, machine.D6, machine.D5, machine.D4, machine.D3, machine.D2)
	keypadDevice.Configure()

	for {
		key := keypadDevice.GetKey()
		if key != keypad4x4.NoKeyPressed {
			println("Button: ", mapping[key])
		}
	}
}
