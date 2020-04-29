package main

import (
	"machine"

	drivers "tinygo.org/x/drivers/hcsr04"
)


func main() {
	// trigPin := machine.Pin(9)
	// echoPin := machine.Pin(10)

	// trigPin.Configure(machine.PinConfig{1})
	// echoPin.Configure(machine.PinConfig{0})
	device := drivers.New(machine.Pin(9), machine.Pin(10))
	device.Configure()
	for {
		println(device.ReadDistance())
	}
}
